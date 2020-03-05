package util

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"go_api_frame/web/common/config"
	"go_api_frame/web/common/database"
)

const DateFormat = "2006-01-02"          //格式化日期
const TimeFormat = "2006-01-02 15:04:05" //格式化时间

func GetPage(page int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * config.MyConfig.PageSize
	}
	return result
}

func Exec(sql string) (bool, error) {
	_, err := database.Engine.Exec(sql)
	if err != nil {
		a := fmt.Sprintf("err:%v,sql:%v", err, sql)
		return false, errors.New(a)
	}
	return true, nil
}

func EncodeMD5(name string) string {
	m := md5.New()
	m.Write([]byte(name))
	return hex.EncodeToString(m.Sum(nil))
}
