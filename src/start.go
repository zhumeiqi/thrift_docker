package main
import (
    "net/http"
    "log"
    "io"
    "os"
    "fmt"
    "os/exec"
    "bytes"
)

func do_thrift(filepath string, language string) {

    cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("./script/thrift_gen %s %s" , language, filepath))

    var out bytes.Buffer
    var errbuf bytes.Buffer

    cmd.Stdout = &out
    cmd.Stderr = &errbuf
    err := cmd.Run()

    if err != nil {
        fmt.Printf("%s %s", out.String())
    }

    fmt.Println("do_thrift", out.String(), errbuf.String())
}

func parse_thrift(w http.ResponseWriter, req * http.Request){

    if req.Method == "POST" {

        language := req.FormValue("language")
        file, handler, err := req.FormFile("uploadfile")

        fmt.Println("language=%s", len(language))

        if err != nil {
            fmt.Println("heeee")
            log.Fatal(err)
            return
        }

        defer file.Close()
        f, err := os.OpenFile("/tmp/"+handler.Filename, os.O_WRONLY | os.O_CREATE, 0666)

        if err != nil{
            log.Fatal(err)
            return
        }
        fmt.Println("Before Copy")

        io.Copy(f, file)
        f.Close()
        fmt.Println("End Copy")

        done := make(chan bool, 1)
        go func(){
           do_thrift("/tmp/"+handler.Filename, language)
           done<-true
        }()

        <-done
        outfile := fmt.Sprintf("/tmp/gen_%s.tar.gz", language)
        http.ServeFile(w, req, outfile)

    }
}

func main(){
    http.HandleFunc("/parse/", parse_thrift)
    err := http.ListenAndServe(":5016", nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
