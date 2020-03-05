package router

import (
	"github.com/gin-gonic/gin"
	"go_api_frame/web/middleware/run"
	"go_api_frame/web/service"
)

func UploadRouter(r *gin.Engine) {
	api := r.Group("/api/v3")
	//上传
	api.POST("/upload",
		run.Run(service.UploadImage))
}
