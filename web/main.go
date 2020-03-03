package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goApiFrame/web/common"
	"goApiFrame/web/middleware/jwt"
	"goApiFrame/web/middleware/log"
	"goApiFrame/web/router"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func init() {
	common.InitConfig()
	common.InitDataEngine()
	log.InitLogger()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-signalChan
		fmt.Println("<<< Cleaning before stop >>>")
		//当收到信号后，会执行相关清理程序或通知各个子进程做自清理。
		//doClean()
		os.Exit(0)
	}()
}

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), jwt.JWT())
	gin.SetMode(common.MyConfig.RunMode)
	router.PactRouter(r)
	router.UserRouter(r)
	err := r.Run(":" + strconv.Itoa(common.MyConfig.Port))
	if err != nil {
		fmt.Println(fmt.Errorf("engine run err %s", err))
	}
}
