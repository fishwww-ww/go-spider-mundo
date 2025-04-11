package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
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
	DB, err = sql.Open("mysql", path)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
		return
	}

	DB.SetConnMaxLifetime(10 * time.Second)
	DB.SetMaxOpenConns(5)

	if err := DB.Ping(); err != nil {
		log.Fatal("数据库 Ping 失败:", err)
		return
	}
	fmt.Println("数据库连接成功")
}
