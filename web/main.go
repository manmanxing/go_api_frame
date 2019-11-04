package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "goApiFrame/web/common"
	"net/http"
)

func init() {
	InitConfig()
	InitLogger()
	InitDataEngine()
}

func main() {
	r := gin.Default()
	r.GET("/test", func(context *gin.Context) {
		context.String(http.StatusOK, "hello goApiFrame")
	})
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("run err:", err)
		Log.Error(err.Error())
	}
}
