package common

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 全局变量
var DB *gorm.DB

// 初始化函数
func InitDB() *gorm.DB {
	dsn := "root:haipeng123@tcp(127.0.0.1:3306)/my-realworld?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect to database:", err)
	}
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)  // 设置最大空闲连接数
	sqlDB.SetMaxOpenConns(100) // 设置最大连接数

	return DB
}

// 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
