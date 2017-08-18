package main

import (
    "bytes"
    "io"
    "fmt"
    "mime/multipart"
    "net/http"
    "os"
    "log"
    "flag"
)

func upload_file(filename string, outputdir string, targetUrl string, params map[string]string) error{

    bodyBuf := new(bytes.Buffer)
    bodyWriter := multipart.NewWriter(bodyBuf)

    fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)

    if err != nil {
        log.Fatal("error writing to buffer")
        return err
    }

    fh, err := os.Open(filename)

    if err != nil {
        log.Fatal("error open file" + filename)
        return err
    }

    defer  fh.Close()

    _, err = io.Copy(fileWriter, fh)

    if err != nil {
        log.Fatal("err copy file")
        return err
    }

    for key, val := range params {
        _ = bodyWriter.WriteField(key, val)
    }

    bodyWriter.Close()

    resp, err := http.Post(targetUrl, bodyWriter.FormDataContentType(), bodyBuf)

    if err != nil {
        return err
    }

    defer resp.Body.Close()

    if resp.Status == "200 OK" {
        file, _ := os.Create(outputdir+"/"+filename+".tar.gz")
        defer file.Close()

        io.Copy(file, resp.Body)
    } else{
        fmt.Println(resp.Status)
    }

    return nil
}


func main() {

    params := make(map[string]string)

    input := flag.String("input", "test.thrift", "输入文件")
    language := flag.String("lang", "go", "生成语言")
    output_path := flag.String("output", "./gen_output",  "输出目录")

    flag.Parse()
    s := fmt.Sprintf("%s %s", *language, *output_path)
    fmt.Println(s)

    params["language"] = *language

    target_url := fmt.Sprintf("http://127.0.0.1:5016/parse/")

    filename := input
    fmt.Println(*filename, target_url, params)

    err := upload_file(*filename, *output_path, target_url, params)

    if err != nil {
        log.Fatal(err)
    }
}


