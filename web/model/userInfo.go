package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"time"
)

type UserInfo struct {
	Id         string `xorm:"varchar(50) 'id'" json:"id" valid:",optional"`
	Name       string `xorm:"varchar(50) 'name'" json:"name" valid:"stringlength(1|50)"`
	Status     int    `xorm:"smallint(2) 'status'" json:"status" valid:"range(0|100)"`
	CustomerId string `xorm:"varchar(50) 'customer_id'" json:"customer_id" valid:"stringlength(1|50)"`
	UserId     string `xorm:"varchar(50) 'user_id'" json:"user_id" valid:"stringlength(1|50)"`
	//	Money       string    `xorm:"decimal(12,2) 'money'" json:"money"`
	Remark      string    `xorm:"varchar(500) 'remark'" json:"remark" valid:"stringlength(0|500)"`
	StartTime   string    `xorm:"DATE 'start_time'" json:"start_time" valid:"data,required"`
	EndTime     string    `xorm:"DATE 'end_time'" json:"end_time" valid:"data,required"`
	CreatedTime time.Time `xorm:"DATETIME created 'create_time'" json:"create_time" valid:"-"`
	UpdatedTime time.Time `xorm:"DATETIME updated 'update_time'" json:"update_time" valid:"-"`
}

func main() {
	u := UserInfo{
		Id:          "",
		Name:        "admin",
		Status:      1,
		CustomerId:  "edeqcfadc",
		UserId:      "sadafadf",
		Remark:      "",
		StartTime:   "",
		EndTime:     "",
		CreatedTime: time.Time{},
		UpdatedTime: time.Time{},
	}
	//全局变量
	//当设置为true时，如果没有定义valid tag，则会提示错误
	//当设置为false时，如果没有定义valid tag，不会提示错误。默认值就是false
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.TagMap["duck"] = func(str string) bool {
		return str == "duck"
	}
	result, err := govalidator.ValidateStruct(u)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
