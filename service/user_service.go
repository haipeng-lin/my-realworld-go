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
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("请检查")))
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

// 用户登录、认证
func UsersLogin(c *gin.Context) {
	// 登录请求结构体
	var LoginUser struct {
		User struct {
			Email    string `json:"email" binding:"required,email"`
			Password string `json:"password" binding:"required"`
		} `json:"user"`
	}

	// 绑定请求数据
	if err := c.ShouldBindJSON(&LoginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请确认数据填写规范"})
		return
	}

	// 获取user
	user := LoginUser.User

	// 根据邮箱和密码查询，判断用户是否存在
	userModel, err := models.SelectUser(&models.UserModel{Email: user.Email, Password: user.Password})
	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("邮箱或密码错误，请检查")))
		return
	}

	//
	userVo := models.UserVo{
		Username: userModel.Username,
		Email:    userModel.Email,
		Bio:      userModel.Bio,
		Image:    userModel.Image,
		Token:    common.GenToken(userModel.ID),
	}

	c.JSON(http.StatusOK, gin.H{"user": userVo})
}

// 关注用户
func FollowUser(c *gin.Context) {
}

// 取消关注用户
func UnFollowUser(c *gin.Context) {
}
