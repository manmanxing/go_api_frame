package common

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var MyConfig Config

//使用yaml，初始化配置文件
type Config struct {
	Port        string `yaml:"port"`
	Connect     string `yaml:"connect"`
	ConnectType string `yaml:"connectType"`

	Loglevel       string `yaml:"loglevel"`
	HookMaxSize    int    `yaml:"hookMaxSize"`
	HookFilename   string `yaml:"hookFilename"`
	HookMaxBackups int    `yaml:"hookMaxBackups"`
	HookMaxAge     int    `yaml:"hookMaxAge"`
	HookCompress   bool   `yaml:"hookCompress"`
	ServiceName    string `yaml:"serviceName"`
}

func InitConfig() {
	config := new(Config)
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("read yaml file err:", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		fmt.Println("unmarshal yaml err:", err)
	}
	MyConfig = *config
}
