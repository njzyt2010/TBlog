package database

import (
	"TBlog/pkg/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init()  {
	dbConf := config.GetDbConf()
	var username,password,host,port,database = dbConf.Username,dbConf.Password,dbConf.Host,dbConf.Port,dbConf.Database

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)

	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err ==nil {
		DB = db
	}
}
