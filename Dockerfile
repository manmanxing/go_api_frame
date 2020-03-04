#拉取 golang 最新的镜像
FROM golang:latest
#设置 golang 环境变量
ENV GOPROXY https://goproxy.cn,direct

#参考网址：https://github.com/manmanxing/go_api_frame
#设置工作目录
WORKDIR $GOPATH/src/github.com/manmanxing/go_api_frame
#将当前上下文目录的内容复制到工作目录中
COPY . $GOPATH/src/github.com/manmanxing/go_api_frame
#执行编译
RUN go build ./web/main.go
#设置暴露端口号
EXPOSE 8000
#启动程序，也就是我们所编译的可执行文件
ENTRYPOINT ["./go_api_frame"]