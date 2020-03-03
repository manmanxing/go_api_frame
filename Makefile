#项目路径为 /home/USER/xxx/work/src/goApiFrame
export GOPATH=/home/USER/xxx/work
APPS=goApiFrame
all: $(APPS)
ad_backend_local:
	go build -o $@  web/main.go
check: test_log
test_log:
	cd $(PWD)/src/common/log && go test
clean:
	rm -fr $(APPS)
