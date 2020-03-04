package model

type UserInfo struct {
	Id       int    `xorm:"int pk autoincr 'id'" json:"id"`
	Username string `xorm:"varchar(50) 'username' notnull "json:"username"`
	Password string `xorm:"varchar(50) 'password' notnull "json:"password"`
}
