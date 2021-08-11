package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"reflect"
	"time"
)

var DB *gorm.DB

func initDB() {
	if true{
		initPG()
	}else{
		initMysql()
	}

	sqlDB, err := DB.DB()
	if err != nil {
		Log.Panic("")
		return
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func initPG() {
	dsn := "host=127.0.0.1 user=dev password=123456 dbname=supo_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		Log.Panic("Connect database error:", err)
	}
}

func initMysql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/dd_web?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		Log.Panic("Connect database error:", err)
	}
}

func autoMigrate(model interface{}) {
	err := DB.AutoMigrate(model)
	if err != nil {
		Log.Errorf("Auto migrate model [%s] error:%s", reflect.ValueOf(model).Type(), err.Error())
	}
}

func expression(domain []interface{}) string {
	var express string
	for _, expr := range domain {
		express += " " + fmt.Sprint(expr)
	}
	return express
	//var express string
	//var values []interface{}
	//var len = len(domain)
	//for i := 0; i <= len; i++ {
	//	expr := domain[i]
	//	switch expr {
	//	case "&":
	//		express+=" and " +
	//	case "|":
	//	case "!":
	//	default:
	//		if express != "" {
	//			express += " and"
	//		}
	//		express += " " + express
	//	}
	//}
}
