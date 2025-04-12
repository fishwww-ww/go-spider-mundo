package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strings"
)

const (
	USERNAME = "root"
	PASSWORD = "123456"
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DBNAME   = "mysql"
)

func InitDB() {
	path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", HOST, ":", PORT, ")/", DBNAME, "?charset=utf8"}, "")
	var err error
	DB, err = gorm.Open(mysql.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	err = DB.AutoMigrate(&Content{})
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}
}
