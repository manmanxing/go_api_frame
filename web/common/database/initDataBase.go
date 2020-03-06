package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"go_api_frame/web/common/config"
	"log"
	"time"
	"xorm.io/core"
)

var MyEngine *xorm.Engine

func InitDataEngine() {
	engine, err := xorm.NewEngine(config.MyConfig.ConnectType, config.MyConfig.Connect)
	if err != nil {
		log.Fatal("init data engine err:", err)
	}
	//在控制台打印出SQL语句
	//engine.ShowSQL(true)
	//设置结构体到表结构的映射关系
	engine.SetMapper(core.GonicMapper{})
	//空闲连接数
	engine.SetMaxIdleConns(config.MyConfig.MaxIdleConns)
	//打开连接数
	engine.SetMaxOpenConns(config.MyConfig.MaxOpenConns)
	//设置最大生命周期
	engine.SetConnMaxLifetime(config.MyConfig.ConnMaxLifetime)
	//设置时区
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
	MyEngine = engine
}
