package model

import (
	"time"
)

type UserInfo struct {
	Id          string    `xorm:"varchar(50) 'id'" json:"id" valid:",optional"`
	Name        string    `xorm:"varchar(50) 'name'" json:"name" valid:"stringlength(1|10)"`
	Password    string    `xorm:"varchar(50) 'password'" json:"password" valid:"stringlength(6|10)"`
	Status      int       `xorm:"smallint(2) 'status'" json:"status" valid:"range(0|100)"`
	Remark      string    `xorm:"varchar(500) 'remark'" json:"remark" valid:",optional"`
	StartTime   string    `xorm:"DATE 'start_time'" json:"start_time" valid:"date,required"`
	EndTime     string    `xorm:"DATE 'end_time'" json:"end_time" valid:"date,required"`
	CreatedTime time.Time `xorm:"DATETIME created 'create_time'" json:"create_time" valid:"-"`
	UpdatedTime time.Time `xorm:"DATETIME updated 'update_time'" json:"update_time" valid:"-"`
}

func (user UserInfo) Create() UserInfo {
	u := UserInfo{
		Id:          "",
		Name:        "admin",
		Password:    "12345",
		Status:      1,
		Remark:      "",
		StartTime:   "2019-01-01",
		EndTime:     "2019-01-02",
		CreatedTime: time.Time{},
		UpdatedTime: time.Time{},
	}
	return u
}
