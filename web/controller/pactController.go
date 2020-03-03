package controller

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"goApiFrame/web/common"
	"goApiFrame/web/errcode"
	"goApiFrame/web/middleware/log"
	"goApiFrame/web/model"
	"strconv"
)

func PactCreate(ctx *gin.Context) interface{} {
	p := new(model.PactInfo)
	valid := validation.Validation{}
	b, err := valid.Valid(p)
	if err != nil {
		log.SugarLogger.Error("err:", err)
	}
	if !b {
		for _, err := range valid.Errors {
			log.SugarLogger.Error("err:", err)
			panic(err)
		}
	}
	return p.Insert()
}

func PactFind(ctx *gin.Context) interface{} {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	start := common.GetPage(page)
	p := new(model.PactInfo)
	return p.Find(start)
}

func PactDelete(ctx *gin.Context) interface{} {
	id, _ := strconv.Atoi(ctx.Param("id"))
	p := new(model.PactInfo)
	return p.Delete(id)
}

func PactUpdate(ctx *gin.Context) interface{} {
	id, _ := strconv.Atoi(ctx.Param("id"))
	p := new(model.PactInfo)
	err := ctx.ShouldBind(p)
	if err != nil {
		log.SugarLogger.Error("err:", err)
		panic(errcode.Database_err)
	}
	//如果只是更新单个字段，那就单独 valid字段
	valid := validation.Validation{}
	b, err := valid.Valid(p)
	if err != nil {
		log.SugarLogger.Error("err:", err)
	}
	if !b {
		for _, err := range valid.Errors {
			log.SugarLogger.Error("err:", err)
			panic(err)
		}
	}
	return p.Update(p.Name, id)
}
