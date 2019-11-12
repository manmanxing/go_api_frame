package controller

import (
	"github.com/gin-gonic/gin"
	"goApiFrame/web/middleware/validator"
	"goApiFrame/web/model"
	"goApiFrame/web/resultInfo"
	"time"
)

func UserCreate(ctx *gin.Context) interface{} {
	var user model.UserInfo
	user = model.UserInfo{
		Id:          "",
		Name:        "admin",
		Password:    "123456",
		Status:      1,
		Remark:      "",
		StartTime:   "2019-01-01",
		EndTime:     "2019-01-02",
		CreatedTime: time.Time{},
		UpdatedTime: time.Time{},
	}
	if !validator.CheckValidator(user) {
		panic(resultInfo.Params_err)
	}
	return user.Create()
}
