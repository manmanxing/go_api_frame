package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"
)

var MyConfig Config

//使用yaml，初始化配置文件
type Config struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	RunMode         string        `yaml:"runMode"`
	Connect         string        `yaml:"connect"`
	ConnectType     string        `yaml:"connectType"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
	PageSize        int           `yaml:"pagesize"`
	Loglevel        string        `yaml:"loglevel"`
	HookMaxSize     int           `yaml:"hookMaxSize"`
	HookMaxBackups  int           `yaml:"hookMaxBackups"`
	HookMaxAge      int           `yaml:"hookMaxAge"`
	HookCompress    bool          `yaml:"hookCompress"`
	ServiceName     string        `yaml:"serviceName"`
	SendEmail       bool          `yaml:"sendEmail"`
	FromEmailUser   string        `yaml:"fromEmailUser"`
	ToEmailUSer     string        `yaml:"toEmailUser"`
	EmailPass       string        `yaml:"emailPass"`
	EmailPort       string        `yaml:"emailPort"`
	EmailSSL        bool          `yaml:"emailSSL"`
	JwtSecret       int           `yaml:"jwtSecret"`
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
