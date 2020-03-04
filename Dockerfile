FROM golang:latest

#拉取golang最新的镜像
ENV GOPROXY https://goproxy.cn,direct

#参考网址：https://github.com/manmanxing/goApiFrame
#设置工作目录
WORKDIR $GOPATH/src/github.com/manmanxing/goApiFrame
#将当前上下文目录的内容复制到工作目录中
COPY . $GOPATH/src/github.com/manmanxing/goApiFrame
#执行编译
RUN go build
#设置暴露端口号
EXPOSE 8000
#启动程序，也就是我们所编译的可执行文件
ENTRYPOINT["./goApiFrame"]