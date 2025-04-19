package routes

import (
	"my-realworld-go/service" 

	"github.com/gin-gonic/gin"
)

// 用户注册登录路由
func Users(router *gin.RouterGroup) {
	router.POST("/", service.UsersRegist)
	router.POST("/login", service.UsersLogin)
}

// 用户资料路由
func Profiles(router *gin.RouterGroup) {
	router.GET("/:username", service.GetProfile)
	router.POST("/:username/follow", service.FollowUser)
	router.DELETE("/:username/follow", service.UnFollowUser)
}
