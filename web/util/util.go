package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goApiFrame/web/common"
	"strconv"
	"time"
)

const DateFormat = "2006-01-02"          //格式化日期
const TimeFormat = "2006-01-02 15:04:05" //格式化时间

func JudgeDate(a ...string) bool {
	if len(a) == 0 {
		return false
	}
	for i := range a {
		_, err := time.Parse(DateFormat, a[i])
		if err != nil {
			return false
		}
	}
	return true
}

func GetPage(c *gin.Context) int {
	result := 0
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		fmt.Println("context query page err:", err)
	}
	if page > 0 {
		result = (page - 1) * common.MyConfig.PageSize
	}
	return result
}
