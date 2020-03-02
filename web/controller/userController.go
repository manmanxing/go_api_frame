package controller

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"goApiFrame/web/model"
	"time"
)

func UserCreate(ctx *gin.Context) interface{} {
	user := &model.PactInfo{
		Id:          "",
		Name:        "admin",
		Status:      1,
		Remark:      "",
		CreatedTime: time.Now().Nanosecond(),
		UpdatedTime: time.Now().Nanosecond(),
		Context:     "",
		ImageUrl:    "",
	}
	valid := validation.Validation{}
	b, err := valid.Valid(user)
	if err != nil {
		panic(err.Error())
	}
	if !b {
		for _, err := range valid.Errors {
			//log.Log.Error("param valid err", zap.Error(err))
			panic(err)
		}
	}
	return user.Create()
}
