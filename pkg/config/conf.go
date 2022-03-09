package config

import (
	"TBlog/pkg/logger"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type configuration struct {
	DbConf     dbConf     `yaml:"db"`
	ServerConf serverConf `yaml:"server"`
}

// DbConf 数据库模型
type dbConf struct {
	Dsn     string `yaml:"dsn"` //数据库连接串
	LogLevel string	`yaml:"log-level"` // 日志级别
}

// serverConf 服务器相关配置
type serverConf struct {
	Port string `yaml:"port"`
}

var conf *configuration
var path string

func init() {
	logger.Info("加载 conf.yml 文件")
	flag.StringVar(&path, "c", "env/conf.yaml", "-c 为配置文件")
	flag.Parse()

	if len(path) <= 0 {
		logger.Panic("配置文件加载失败，请确认配置文件路径是否正确")
		return
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Panic(err)
		return
	}
	var configuration configuration
	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		logger.Panic(err)
		return
	}
	conf = &configuration

}

func GetDbConf() dbConf {
	return conf.DbConf
}
func GetServerConf() serverConf {
	return conf.ServerConf
}
