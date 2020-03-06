package run

import (
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go_api_frame/web/common/config"
	"go_api_frame/web/common/errcode"
	"go_api_frame/web/common/redis"
	"go_api_frame/web/middleware/log"
	"net/http"
)

type handlerFunc func(*gin.Context) interface{}

type Options struct {
	Cache bool //是否缓存
}

func Run(handleFunc handlerFunc, options ...Options) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ErrHandle(context, err)
			}
		}()
		cache := false
		for _, v := range options {
			cache = v.Cache
		}
		if cache {
			if redis.ExistCache(context.Request.URL.Path) {
				cacheValue, err := redis.GetCache(context.Request.URL.Path)
				if err != nil {
					log.SugarLogger.Error("get cache err:", err)
					context.JSON(http.StatusOK, gin.H{
						"code":  errcode.GetCacheFail,
						"error": errcode.GetErr(errcode.GetCacheFail),
						"data":  nil,
					})
					return
				}

				var data interface{}
				err = json.Unmarshal([]byte(cacheValue), &data)
				if err != nil {
					log.SugarLogger.Error("json unmarshal err:", err)
					context.JSON(http.StatusOK, gin.H{
						"code":  errcode.JsonUnmarshalFail,
						"error": errcode.GetErr(errcode.JsonUnmarshalFail),
						"data":  nil,
					})
					return
				}
				context.JSON(http.StatusOK, gin.H{
					"code":  "0",
					"error": nil,
					"data":  data,
				})
				return
			} else {
				result := handleFunc(context)
				cacheByte, _ := json.Marshal(result)
				err := redis.SetCache(context.Request.URL.Path, string(cacheByte), config.MyConfig.RedisExpireTime)
				if err != nil {
					log.SugarLogger.Error("set cache err:", err)
					context.JSON(http.StatusOK, gin.H{
						"code":  errcode.SetCacheFail,
						"error": errcode.GetErr(errcode.SetCacheFail),
						"data":  nil,
					})
					return
				}
				context.JSON(http.StatusOK, gin.H{
					"code":  "0",
					"error": nil,
					"data":  result,
				})
				return
			}
		}

		result := handleFunc(context)
		context.JSON(http.StatusOK, gin.H{
			"code":  "0",
			"error": nil,
			"data":  result,
		})
	}
}

//有两种错误，一种是panic（errcode）,为 string 类型，一种是 validation 的自定义验证错误
func ErrHandle(c *gin.Context, e interface{}) {
	switch value := e.(type) {
	case string:
		result := errcode.GetErr(e.(string))
		//go email.Email(result.Msg, c.Request.Method+"  "+c.Request.Host+c.Request.RequestURI, c.Request.UserAgent(), c.ClientIP())
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  result.Code,
			"error": result.Msg,
			"data":  nil,
		})
	case *validation.Error:
		//go email.Email(value.Msg, c.Request.Method+"  "+c.Request.Host+c.Request.RequestURI, c.Request.UserAgent(), c.ClientIP())
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  errcode.GetErr(errcode.ParamsErr).Code,
			"error": value.Message,
			"data":  nil,
		})
	}
}
