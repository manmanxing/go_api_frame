#项目路径为 /home/USER/xxx/work/src/goApiFrame
export GOPATH=/home/zengjie/work
APPS=goApiFrame
all: $(APPS)
goApiFrame:
	go build -o $@  web/main.go
clean:
	rm -fr $(APPS)
