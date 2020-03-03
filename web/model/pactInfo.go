package model

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"goApiFrame/web/common"
	"goApiFrame/web/common/errcode"
	"goApiFrame/web/common/util"
	"goApiFrame/web/middleware/log"
	"strconv"
	"strings"
)

//合同表
type PactInfo struct {
	Id          int    `xorm:"int pk autoincr 'id'" json:"id"`
	Name        string `xorm:"varchar(50) 'name' notnull  COMMENT '合同名称'"  json:"name" valid:"MaxSize(10);MinSize(1)"` //合同名称
	Status      int    `xorm:"smallint(2) 'status' default(0) COMMENT '合同状态，-1为删除'" json:"status" valid:"Range(-1,2)"` //合同状态
	Context     string `xorm:"text 'context' default('') COMMENT '合同内容'" json:"context"`                               //合同内容
	ImageUrl    string `xorm:"varchar(255) 'image_url' default('') COMMENT '合同照片'" json:"image_url"`                   //合同照片
	Remark      string `xorm:"varchar(255) 'remark' default('') COMMENT '合同备注'" json:"remark" valid:"MaxSize(500)"`    //合同备注
	CreatedTime int    `xorm:"int(10)  'create_time'" json:"create_time" `
	UpdatedTime int    `xorm:"int(10)  'update_time'" json:"update_time" `
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *PactInfo) Valid(v *validation.Validation) {
	if strings.Index(u.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		_ = v.SetError("Name", "名称里不能含有 admin")
	}
}

func (p *PactInfo) Insert() bool {
	_, err := common.Engine.Insert(p)
	if err != nil {
		fmt.Println("insert pact err:", err)
		log.SugarLogger.Error("err:", err)
		panic(errcode.Database_err)
	}
	return true
}

func (p *PactInfo) Find(pageSize int) []PactInfo {
	result := make([]PactInfo, 0)
	err := common.Engine.Limit(common.MyConfig.PageSize, pageSize).OrderBy("create_time").Find(result)
	if err != nil {
		log.SugarLogger.Error("err:", err)
		panic(errcode.Database_err)
	}
	return result
}

func (p *PactInfo) Delete(id int) bool {
	sql := fmt.Sprintf("update pact_info set status = %s where id = %s", strconv.Itoa(util.Delete), strconv.Itoa(id))
	ok, err := util.Exec(sql)
	if err != nil && ok == false {
		log.SugarLogger.Error("err:", err)
		panic(errcode.Params_err)
	}
	return true
}

func (p *PactInfo) Update(name string, id int) bool {
	sql := fmt.Sprintf("update pact_info set name = '%s' where id = %s", name, strconv.Itoa(id))
	ok, err := util.Exec(sql)
	if err != nil && ok == false {
		log.SugarLogger.Error("err:", err)
		panic(errcode.Params_err)
	}
	return true
}
