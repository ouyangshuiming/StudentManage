package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func init() {
	connect, err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/StudentManage")

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("数据库连接成功")
	db = connect
}

func GetDB() *sql.DB {
	return db
}
