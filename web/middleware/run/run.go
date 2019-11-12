package run

import (
	"github.com/gin-gonic/gin"
	"goApiFrame/web/middleware/email"
	"goApiFrame/web/resultInfo"
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

func ErrHandle(c *gin.Context, e interface{}) {
	switch value := e.(type) {
	case string:
		email.Email(e.(string), c.Request.Method+"  "+c.Request.Host+c.Request.RequestURI, c.Request.UserAgent(), c.ClientIP())
		result := resultInfo.GetErr(e.(string))
		c.JSON(500, gin.H{
			"code":  result.Code,
			"error": result.Msg,
			"data":  nil,
		})
	case resultInfo.ResultInfo:
		email.Email(value.Msg, c.Request.Method+"  "+c.Request.Host+c.Request.RequestURI, c.Request.UserAgent(), c.ClientIP())
		c.JSON(500, gin.H{
			"code":  value.Code,
			"error": value.Msg,
			"data":  nil,
		})
	}
}