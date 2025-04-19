package main

import (
	"my-realworld-go/common"
	"my-realworld-go/user" // user包

	"github.com/gin-gonic/gin"         // Gin框架
	_ "github.com/go-sql-driver/mysql" // 加载MySQL驱动
)

func main() {

	// 初始化数据库
	common.InitDB()

	// 初始化路由
	router := gin.Default()

	// api 分组
	apiGroup := router.Group("/api")

	// profiles 分组
	user.Profiles(apiGroup.Group("/profiles"))

	// users 分组
	user.Users(apiGroup.Group("/users"))

	router.Run(":8080")
}
