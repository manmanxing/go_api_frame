package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(context *gin.Context) {
		context.String(http.StatusOK, "hello goApiFrame")
	})
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("run err:", err)
	}
}
