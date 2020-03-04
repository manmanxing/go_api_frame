package main

import (
	"fmt"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	//"github.com/fvbock/endless"
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

// @title Go-site Example API
// @version 1.0
// @description This is a sample go-API server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1
// @BasePath ""
func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), jwt.JWT())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	gin.SetMode(common.MyConfig.RunMode)
	router.PactRouter(r)
	router.UserRouter(r)
	err := r.Run(":" + strconv.Itoa(common.MyConfig.Port))
	if err != nil {
		fmt.Println(fmt.Errorf("engine run err %s", err))
	}
}

//提供重启操作
//- 不关闭现有连接（正在运行中的程序）
//- 新的进程启动并替代旧进程，并接管新的连接
//- 连接要随时响应用户的请求，当用户仍在请求旧进程时要保持连接，新用户应请求新进程，不可以出现拒绝请求的情况
//func main(){
//	r := gin.New()
//	r.Use(gin.Logger(), gin.Recovery(), jwt.JWT())
//	gin.SetMode(common.MyConfig.RunMode)
//	router.PactRouter(r)
//	router.UserRouter(r)
//	endless.DefaultReadTimeOut = common.MyConfig.ReadTimeout
//	endless.DefaultWriteTimeOut = common.MyConfig.WriteTimeout
//	endless.DefaultMaxHeaderBytes = 1 << 20
//	endPoint := fmt.Sprintf(":%s", common.MyConfig.Port)
//	server := endless.NewServer(endPoint,r)
//	//在 BeforeBegin 时输出当前进程的 pid，调用 ListenAndServe 将实际“启动”服务
//	server.BeforeBegin = func(add string) {
//		fmt.Printf("actual pid is %d",syscall.Getpid())
//	}
//	err := server.ListenAndServe()
//	if err != nil {
//		fmt.Printf("server start err:%v",err)
//	}else {
//		fmt.Println("server start success")
//	}
//}
