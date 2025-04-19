package user

import (
	"github.com/gin-gonic/gin"
)

// 用户注册登录——路由
func Users(router *gin.RouterGroup) {
	router.POST("/", UsersRegist)
	router.POST("/login", UsersLogin)
}

// 用户资料——路由
func Profiles(router *gin.RouterGroup) {
	router.GET("/:username", GetProfile)
	router.POST("/:username/follow", FollowUser)
	router.DELETE("/:username/follow", UnFollowUser)
}
