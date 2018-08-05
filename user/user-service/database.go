package main

import (
	"github.com/gohouse/gorose"        //import Gorose
	_ "github.com/go-sql-driver/mysql" //import DB driver
	"os"
)

func CreateConnection() (gorose.Connection, error) {

	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	prefix := os.Getenv("DB_PREFIX")
	// DB Config.(Recommend to use configuration file to import)
	var DbConfig = map[string]interface{}{
		"Default": "mysql_dev",  //数据库默认配置
		"SetMaxOpenConns": 300,  //(连接池)最大打开的连接数,默认值为0表示不限制
		"SetMaxIdleConns": 10,   //(连接池)限制的连接数,默认为1

		// Define the database configuration character "mysql_dev".
		"Connections": map[string]map[string]string{
			"mysql_dev": map[string]string{     // 定义名为 mysql_dev 的数据库配置
				"host":     host,               //地址
				"username": username,			//用户名
				"password": password,			//密码
				"port":     "3306",				//端口
				"database": database,			//数据库名称
				"charset":  "utf8",				//字符集
				"protocol": "tcp",				//协议
				"prefix":   prefix,  			//表前缀
				"driver":   "mysql", 			//数据库驱动(mysql,sqlite,postgres,oracle,mssql)
			},
		},
	}
	// 初始化数据库链接, 默认会链接配置中 default 指定的值
	return gorose.Open(DbConfig)

}
