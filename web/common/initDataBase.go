package common

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	"xorm.io/core"
)

var Engine *xorm.Engine

func InitDataEngine() {
	Engine, err := xorm.NewEngine(MyConfig.ConnectType, MyConfig.Connect)
	if err != nil {
		fmt.Println("init data engine err:", err)
		Log.Error(err.Error())
	}
	//在控制台打印出SQL语句
	Engine.ShowSQL(true)
	//设置结构体到表结构的映射关系
	Engine.SetMapper(core.GonicMapper{})
	//空闲连接数
	Engine.SetMaxIdleConns(MyConfig.MaxIdleConns)
	//打开连接数
	Engine.SetMaxOpenConns(MyConfig.MaxOpenConns)
	//设置最大生命周期
	Engine.SetConnMaxLifetime(MyConfig.ConnMaxLifetime)
	//设置时区
	Engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
}
