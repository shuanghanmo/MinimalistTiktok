package dao

import (
	_ "database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func init() {
	// MySQL 配置信息
	username := "root" // 账号
	//password := "swust6136" // 密码
	password := "123456" // 密码
	host := "127.0.0.1"  // 地址
	port := 33306        // 端口
	DBname := "douyin"   // 数据库名称
	timeout := "10s"     // 连接超时，10秒
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, DBname, timeout)
	// Open 连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("failed to connect mysql.", err)
	}
	DB = db
}
