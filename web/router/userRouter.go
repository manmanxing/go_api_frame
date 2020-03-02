package router

import (
	"github.com/gin-gonic/gin"
	"goApiFrame/web/controller"
	"goApiFrame/web/middleware/run"
)

func UserRouter(r *gin.Engine) {
	api := r.Group("/api/v1")
	//获取用户列表
	api.GET("/user",
		run.Run(controller.UserCreate))
	//新增用户
	api.POST("/user",
		run.Run(controller.UserCreate))
	//更新指定的用户信息
	api.PUT("/user/:id",
		run.Run(controller.UserCreate))
	//删除指定的用户
	api.DELETE("/user/:id",
		run.Run(controller.UserCreate))
}
