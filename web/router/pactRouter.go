package router

import (
	"github.com/gin-gonic/gin"
	"goApiFrame/web/controller"
	"goApiFrame/web/middleware/run"
)

func PactRouter(r *gin.Engine) {
	api := r.Group("/api/v1")
	//获取合同列表
	api.GET("/pact",
		run.Run(controller.PactFind))
	//新增合同
	api.POST("/pact",
		run.Run(controller.PactCreate))
	//更新指定的合同信息
	api.PUT("/pact/:id",
		run.Run(controller.PactUpdate))
	//删除指定的合同
	api.DELETE("/pact/:id",
		run.Run(controller.PactDelete))
}
