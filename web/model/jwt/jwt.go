package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"goApiFrame/web/common"
	"goApiFrame/web/model"
	"time"
)

var jwtSecret []byte

//自定义payload结构体
type Claims struct {
	UserInfo *model.UserInfo
	jwt.StandardClaims
}

//实现 `type Claims interface` 的 `Valid() error` 方法,自定义校验内容
//func (c Claims) Valid() (err error) {
//	//验证token是否过期
//	//如果exp为0 返回 ！req
//	//如果返回 false，则表示过期
//	if c.VerifyExpiresAt(time.Now().Unix(), true) == false {
//		return errors.New("token is expired")
//	}
//	if !c.VerifyIssuer(MyConfig.ServiceName, true) {
//		return errors.New("token's issuer is wrong")
//	}
//	if len(c.UserInfo.Username) == 0 {
//		return errors.New("invalid username in jwt")
//	}
//	return
//}

//生成JWT-string
func GenerateToken(m *model.UserInfo) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour) //过期时间

	claims := Claims{
		UserInfo: m,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    common.MyConfig.ServiceName, //服务名,
		},
	}
	//包含SigningMethodHS256、SigningMethodHS384、SigningMethodHS512三种crypto.Hash方案
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

//解析 payload 内容，并获取用户信息
func ParseToken(token string) (*Claims, error) {
	if token == "" {
		return nil, errors.New("token len = 0")
	}
	// 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
