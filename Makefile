#项目路径为 /home/USER/xxx/work/src/goApiFrame
export GOPATH=/home/zengjie/work
APPS=goApiFrame
all: $(APPS)
goApiFrame:
	go build -o $@  web/main.go
check: test_log
test_log:
	cd $(PWD)/src/common/log && go test
clean:
	rm -fr $(APPS)
