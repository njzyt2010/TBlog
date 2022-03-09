package database

import (
	"TBlog/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
)

var DB *gorm.DB

func init() {
	dbConf := config.GetDbConf()
	// 将配置文件中的参数转为小写
	lowerLogLevel := strings.ToLower(dbConf.LogLevel)
	logLevel := convertLogLevel(lowerLogLevel)

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	db, err := gorm.Open(mysql.Open(dbConf.Dsn), config)
	if err == nil {
		DB = db
	}
}

// 日志级别转换
func convertLogLevel(loglevel string) logger.LogLevel {
	var retLevel = logger.Error
	switch loglevel {
	case "info":
		retLevel = logger.Info
	case "warn":
		retLevel = logger.Warn
	case "error":
		retLevel = logger.Error
	case "silent":
		retLevel = logger.Silent
	default:
		retLevel = logger.Error
	}
	return retLevel
}
