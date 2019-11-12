package controller

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
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
			//log.Println(err.Key, err.Message)
			panic(err)
		}
	}
	return user.Create()
}
