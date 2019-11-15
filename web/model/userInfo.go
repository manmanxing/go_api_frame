package model

import (
	"github.com/astaxie/beego/validation"
	"goApiFrame/web/util"
	"strings"
	"time"
)

type UserInfo struct {
	Id          string    `xorm:"varchar(50) 'id'" json:"id"`
	Name        string    `xorm:"varchar(50) 'name'" json:"name" valid:"MaxSize(10);MinSize(1)"`
	Password    string    `xorm:"varchar(50) 'password'" json:"password" valid:"MaxSize(10);MinSize(6)"`
	Status      int       `xorm:"smallint(2) 'status'" json:"status" valid:"Range(0,2)"`
	Remark      string    `xorm:"varchar(500) 'remark'" json:"remark" valid:"MaxSize(500)"`
	StartTime   string    `xorm:"DATE 'start_time'" json:"start_time"`
	EndTime     string    `xorm:"DATE 'end_time'" json:"end_time"`
	CreatedTime time.Time `xorm:"DATETIME created 'create_time'" json:"create_time" `
	UpdatedTime time.Time `xorm:"DATETIME updated 'update_time'" json:"update_time" `
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *UserInfo) Valid(v *validation.Validation) {
	if strings.Index(u.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		_ = v.SetError("Name", "名称里不能含有 admin")
	}
	if util.JudgeDate(u.StartTime, u.EndTime) {
	} else {
		_ = v.SetError("Date", "日期格式错误")
	}
}

func (user UserInfo) Create() UserInfo {
	return user
}
