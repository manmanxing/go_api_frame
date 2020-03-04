package controller

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"goApiFrame/web/common/errcode"
	"goApiFrame/web/common/util"
	"goApiFrame/web/middleware/log"
	"goApiFrame/web/model"
	"strconv"
)

// @Summary 新增合同
// @Accept  json
// @Produce  json
// @Param name query string true "name"
// @Param context query string true "context"
// @Param image_url query string true "image_url"
// @Param remark query string true "remark"
// @Success 200 {string} json "{"code":"0","data":true,"error":null}"
// @Failure 500 {string} json "{"code":"错误码","data":null,"error":"错误信息"}"
// @Router /api/v1/pact [post]
func PactCreate(ctx *gin.Context) interface{} {
	p := new(model.PactInfo)
	err := ctx.ShouldBind(p)
	if err != nil {
		log.SugarLogger.Error("err:", err)
		panic(errcode.Database_err)
	}
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

// @Summary 查找合同列表
// @Produce  json
// @Param page path int false "page"
// @Success 200 {string} json "{"code":"0","data":{},"error":null}"
// @Failure 500 {string} json "{"code":"错误码","data":null,"error":"错误信息"}"
// @Router /api/v1/pact [get]
func PactFind(ctx *gin.Context) interface{} {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	start := util.GetPage(page)
	p := new(model.PactInfo)
	return p.Find(start)
}

// @Summary 删除合同
// @Produce  json
// @Param id path int true "id"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 500 {string} json "{"code":"错误码","data":null,"error":"错误信息"}"
// @Router /api/v1/pact/{id} [delete]
func PactDelete(ctx *gin.Context) interface{} {
	id, _ := strconv.Atoi(ctx.Param("id"))
	p := new(model.PactInfo)
	return p.Delete(id)
}

// @Summary 更新合同名称
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Param name query string true "name"
// @Success 200 {string} json "{"code":"0","data":true,"error":null}"
// @Failure 500 {string} json "{"code":"错误码","data":null,"error":"错误信息"}"
// @Router /api/v1/pact/{id} [put]
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
