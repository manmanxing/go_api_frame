#项目路径为 /home/USER/xxx/work/src/goApiFrame
export GOPATH=/home/zengjie/work
APPS=go_api_frame
all: $(APPS)
go_api_frame:
	go build -o $@  web/main.go
clean:
	rm -fr $(APPS)
