package common

import (
	"goApiFrame/web/errcode"
	"goApiFrame/web/middleware/log"
	"time"
)

const DateFormat = "2006-01-02"          //格式化日期
const TimeFormat = "2006-01-02 15:04:05" //格式化时间
const Delete = -1

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

func GetPage(page int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * MyConfig.PageSize
	}
	return result
}

func Exec(sql string) bool {
	_, err := Engine.Exec(sql)
	if err != nil {
		log.SugarLogger.Error("err:", err, " sql:", sql)
		panic(errcode.Database_err)
	}
	return true
}
