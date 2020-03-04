package router

import (
	"github.com/gin-gonic/gin"
	"go_api_frame/web/controller"
	"go_api_frame/web/middleware/run"
)

func UserRouter(r *gin.Engine) {
	api := r.Group("/api/v2")
	//获取用户列表
	api.GET("/user",
		run.Run(controller.UserCreate))
}
