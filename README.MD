## 快速配置thrift编译环境
使用Docker启动容器来编译thrift协议文件，docker镜像自动配置好了thrift的编译环境
client.go 通过http的方式上传需要编译的文件到docker中，服务器将编译好的生成文件发回给客户端
### 启动方式
    cd ${ROOT}/docker
    docker-compose up
### 编译方式
    go run client.go --lang=py
    go run client.go --lang=go
## 缺陷
    目前只支持编译单文件，方便初步学习thrift，如果要支持多文件可以扩展
