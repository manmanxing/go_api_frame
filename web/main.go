package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goApiFrame/web/common"
	"goApiFrame/web/router"
	"strconv"
)

func init() {
	common.InitConfig()
	common.InitDataEngine()
	//InitLogger()
}

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(common.MyConfig.RunMode)
	router.UserRouter(r)
	err := r.Run(":" + strconv.Itoa(common.MyConfig.Port))
	if err != nil {
		fmt.Println(fmt.Errorf("engine run err %s", err))
	}
}
