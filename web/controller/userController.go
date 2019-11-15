package controller

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"goApiFrame/web/middleware/log"
	"goApiFrame/web/model"
	"time"
)

func UserCreate(ctx *gin.Context) interface{} {
	user := &model.UserInfo{
		Id:          "",
		Name:        "admin",
		Password:    "123456",
		Status:      1,
		Remark:      "",
		StartTime:   "2019-01-ergfer01",
		EndTime:     "2019-01-02",
		CreatedTime: time.Time{},
		UpdatedTime: time.Time{},
	}
	valid := validation.Validation{}
	b, err := valid.Valid(user)
	if err != nil {
		panic(err.Error())
	}
	if !b {
		for _, err := range valid.Errors {
			log.Log.Error("param valid err", zap.Error(err))
			panic(err)
		}
	}
	return user.Create()
}
