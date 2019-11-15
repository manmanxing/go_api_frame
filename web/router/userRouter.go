package router

import (
	"github.com/gin-gonic/gin"
	"goApiFrame/web/controller"
	"goApiFrame/web/middleware/run"
)

func UserRouter(r *gin.Engine) {
	r.GET("/test/user",
		run.Run(controller.UserCreate))
}
