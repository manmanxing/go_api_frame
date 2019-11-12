package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "goApiFrame/web/common"
	"goApiFrame/web/middleware/log"
	"goApiFrame/web/middleware/validator"
	"goApiFrame/web/router"
	"net/http"
)

func init() {
	InitConfig()
	InitDataEngine()
}

func main() {
	//gin.SetMode(gin.ReleaseMode)  //生产环境使用
	r := gin.Default()
	r.Use(log.Logger(), validator.Validator())
	r.GET("/test", func(context *gin.Context) {
		context.String(http.StatusOK, "hello goApiFrame")
	})
	router.UserRouter(r)
	err := r.Run(":" + MyConfig.Port)
	if err != nil {
		fmt.Println(fmt.Errorf("engine run err %s", err))
	}
}
