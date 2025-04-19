package service

import (
	"errors"
	"my-realworld-go/common"
	"my-realworld-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取个人资料
func GetProfile(c *gin.Context) {
	username := c.Param("username")
	// 查询 User
	userModel, err := models.SelectUser(&models.UserModel{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("非法")))
		return
	}
	// 转换为 Profile 对象
	profile := models.Profile{
		Username: userModel.Username,
		Bio:      userModel.Bio,
		Image:    userModel.Image,
		// TODO：先默认为false，待改
		Following: false,
	}

	c.JSON(http.StatusOK, gin.H{"profile": profile})
}

// 用户注册
func UsersRegist(c *gin.Context) {
}

// 用户登录
func UsersLogin(c *gin.Context) {
}

// 关注用户
func FollowUser(c *gin.Context) {
}

// 取消关注用户
func UnFollowUser(c *gin.Context) {
}
