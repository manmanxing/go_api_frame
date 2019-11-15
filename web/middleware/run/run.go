package run

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"goApiFrame/web/errcode"
)

type handlerFunc func(*gin.Context) interface{}

func Run(handleFunc handlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		//处理异常
		defer func() {
			if err := recover(); err != nil {
				ErrHandle(context, err)
			}
		}()
		result := handleFunc(context)
		context.JSON(200, gin.H{
			"code":  "0",
			"error": nil,
			"data":  result,
		})
	}
}

//有两种错误，一种是panic（errcode），一种是 validation 的自定义验证错误
func ErrHandle(c *gin.Context, e interface{}) {
	switch value := e.(type) {
	case string:
		result := errcode.GetErr(e.(string))
		//go email.Email(result.Msg, c.Request.Method+"  "+c.Request.Host+c.Request.RequestURI, c.Request.UserAgent(), c.ClientIP())
		c.JSON(500, gin.H{
			"code":  result.Code,
			"error": result.Msg,
			"data":  nil,
		})
	case *validation.Error:
		//go email.Email(value.Msg, c.Request.Method+"  "+c.Request.Host+c.Request.RequestURI, c.Request.UserAgent(), c.ClientIP())
		c.JSON(500, gin.H{
			"code":  "-1",
			"error": value.Message,
			"data":  nil,
		})
	}
}
