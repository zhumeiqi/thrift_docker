## 说明
学习thrift过程中安装编译环境比较麻烦，自己也正好要学习Go语言及Docker等相关内容，就顺便搞了这个小项目
学习性质，后续有需要再扩展
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
