package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "goApiFrame/web/common"
	. "goApiFrame/web/middleware/log"
	"goApiFrame/web/router"
	"strconv"
)

func init() {
	InitConfig()
	InitDataEngine()
	InitLogger()
}

func main() {
	//gin.SetMode(gin.ReleaseMode)  //生产环境使用
	r := gin.Default()
	r.Use(Logger())
	router.TestUserRouter(r)
	err := r.Run(":" + strconv.Itoa(MyConfig.Port))
	if err != nil {
		fmt.Println(fmt.Errorf("engine run err %s", err))
	}
}
