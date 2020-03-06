package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

var MyConfig Config

//使用yaml，初始化配置文件
type Config struct {
	//数据库配置
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	Connect         string        `yaml:"connect"`
	ConnectType     string        `yaml:"connectType"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
	//运行模式
	RunMode string `yaml:"runMode"`
	//分页
	PageSize int `yaml:"pagesize"`
	//日志
	Loglevel       string `yaml:"loglevel"`
	HookMaxSize    int    `yaml:"hookMaxSize"`
	HookMaxBackups int    `yaml:"hookMaxBackups"`
	HookMaxAge     int    `yaml:"hookMaxAge"`
	HookCompress   bool   `yaml:"hookCompress"`
	LogSavePath    string `yaml:"logSavePath"`
	//服务器配置
	ServiceName string `yaml:"serviceName"`
	//email配置
	SendEmail     bool   `yaml:"sendEmail"`
	FromEmailUser string `yaml:"fromEmailUser"`
	ToEmailUSer   string `yaml:"toEmailUser"`
	EmailPass     string `yaml:"emailPass"`
	EmailPort     string `yaml:"emailPort"`
	EmailSSL      bool   `yaml:"emailSSL"`
	//图片上传配置
	ImageSavePath  string `yaml:"imageSavePath"`
	ImageMaxSize   int64  `yaml:"imageMaxSize"`
	ImageAllowExts string `yaml:"imageAllowExts"`
	//redis
	RedisHost        string        `yaml:"redisHost"`
	RedisPwd         string        `yaml:"redisPwd"`
	RedisMaxIde      int           `yaml:"redisMaxIde"`      //最大空闲连接数
	RedisActive      int           `yaml:"redisActive"`      //在给定时间内，允许分配的最大连接数（当为零时，没有限制）
	RedisExpireTime  int           `yaml:"redisExpireTime"`  //过期时间
	RedisIdleTimeout time.Duration `yaml:"redisIdleTimeout"` //在给定时间内将会保持空闲状态，若到达时间限制则关闭连接（当为零时，没有限制）
}

func InitConfig() {
	config := new(Config)
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		log.Fatal("read yaml file err:", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatal("unmarshal yaml err:", err)
	}
	MyConfig = *config
}
