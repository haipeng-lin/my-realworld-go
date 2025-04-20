package user

import (
	"github.com/gin-gonic/gin"
)

// 用户注册、登录——路由
func UsersRoute(router *gin.RouterGroup) {
	router.POST("/", UsersRegist)
	router.POST("/login", UsersLogin)
}

// 用户修改、获取——路由
func UserRoute(router *gin.RouterGroup) {
	router.GET("/", GetUser)
	router.PUT("/", UserUpdate)
}

// 用户资料——路由
func ProfilesRoute(router *gin.RouterGroup) {
	router.GET("/:username", GetProfile)
	router.POST("/:username/follow", FollowUser)
	router.DELETE("/:username/follow", UnFollowUser)
}
