package validator

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"goApiFrame/web/util"
	"time"
)

func Validator() gin.HandlerFunc {
	//全局变量
	//当设置为true时，如果没有定义valid tag，则会提示错误
	//当设置为false时，如果没有定义valid tag，不会提示错误。默认值就是false
	//字符串使用utf8.RuneCountInString统计长度
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.TagMap["date"] = func(str string) bool {
		_, err := time.Parse(util.DateFormat, str)
		return err == nil
	}
	return func(context *gin.Context) {
		context.Next()
	}
}

func CheckValidator(object interface{}) bool {
	result, err := govalidator.ValidateStruct(object)
	if err != nil {
		fmt.Println(err)
	}
	return result
}
