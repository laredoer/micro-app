package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"fmt"
)

func CreateConnection() (*gorm.DB,error){
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return gorm.Open("mysql",fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8",
				user,password,host,dbName,
		),
	)
}
