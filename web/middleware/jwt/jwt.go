package jwt

import (
	"github.com/gin-gonic/gin"
	"goApiFrame/web/common"
	"goApiFrame/web/errcode"
	"net/http"
	"time"
)

const contextKeyUserObj = "authedUserObj"

//从 url-query 的 token 获取JWT-string
//model.JwtParseUser(token)解析JWT-string获取User结构体(减少中间件查询数据库的操作和时间)
//设置用户信息到gin.Context 其他的handler通过gin.Context.Get(contextKeyUserObj),在进行用户Type Assert得到model.User 结构体.
//使用了jwt-middle之后的handle从gin.Context中获取用户信息
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := ""
		var data interface{}
		token := c.Query("token")
		if token == "" {
			code = errcode.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else {
			claims, err := common.ParseToken(token)
			if err != nil {
				code = errcode.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = errcode.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			} else {
				c.Set(contextKeyUserObj, claims.UserInfo)
			}
		}

		if code != "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  errcode.GetErr(code),
				"data": data,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
