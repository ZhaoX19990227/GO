package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/fmmall2?charset=utf8")
	if err != nil {
		log.Fatal("连接错误", err)
	} else {
		fmt.Println("连接成功", db)
	}
}
