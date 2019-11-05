package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"goApiFrame/web/util"
	"time"
)

type UserInfo struct {
	Id          string    `xorm:"varchar(50) 'id'" json:"id" valid:",optional"`
	Name        string    `xorm:"varchar(50) 'name'" json:"name" valid:"stringlength(1|50)"`
	Password    string    `xorm:"varchar(50) 'password'" json:"password" valid:"stringlength(1|50)"`
	Status      int       `xorm:"smallint(2) 'status'" json:"status" valid:"range(0|100)"`
	Remark      string    `xorm:"varchar(500) 'remark'" json:"remark" valid:",optional"`
	StartTime   string    `xorm:"DATE 'start_time'" json:"start_time" valid:"date,required"`
	EndTime     string    `xorm:"DATE 'end_time'" json:"end_time" valid:"date,required"`
	CreatedTime time.Time `xorm:"DATETIME created 'create_time'" json:"create_time" valid:"-"`
	UpdatedTime time.Time `xorm:"DATETIME updated 'update_time'" json:"update_time" valid:"-"`
}

func main() {
	u := UserInfo{
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
	//全局变量
	//当设置为true时，如果没有定义valid tag，则会提示错误
	//当设置为false时，如果没有定义valid tag，不会提示错误。默认值就是false
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.TagMap["date"] = func(str string) bool {
		_, err := time.Parse(util.DateFormat, str)
		return err == nil
	}
	result, err := govalidator.ValidateStruct(u)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
