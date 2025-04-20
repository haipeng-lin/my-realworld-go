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

	// Users 分组：登录注册（不用身份验证）
	user.UsersRoute(apiGroup.Group("/users"))

	// jwt——身份验证
	apiGroup.Use(user.AuthMiddleware())

	// profiles 、user 分组（用身份验证）
	user.ProfilesRoute(apiGroup.Group("/profiles"))
	user.UserRoute(apiGroup.Group("/user"))

	router.Run(":8080")
}
