FROM thrift 

RUN sed -i 's/httpredir.debian.org/mirrors.aliyun.com/' /etc/apt/sources.list
RUN sed -i 's/deb-src/# deb-src/' /etc/apt/sources.list
RUN apt-get update && apt-get -y install golang

RUN chmod a+x /etc/run/start_go_server

CMD ["/etc/run/start_go_server"]

