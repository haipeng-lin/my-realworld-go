package user

import (
	"errors"
	"my-realworld-go/common"

	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取个人资料
func GetProfile(c *gin.Context) {
	username := c.Param("username")
	// 查询 User
	userModel, err := SelectUser(&UserModel{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("请检查")))
		return
	}
	// 转换为 Profile 对象
	profile := Profile{
		Username: userModel.Username,
		Bio:      userModel.Bio,
		Image:    userModel.Image,
		// TODO：先默认为false，待改
		Following: false,
	}

	c.JSON(http.StatusOK, gin.H{"profile": profile})
}

// 用户登录、认证
func UsersLogin(c *gin.Context) {

	// 绑定请求数据
	if err := c.ShouldBindJSON(&LoginUserDTO); err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("登录", errors.New("请确认数据填写是否规范！")))
		return
	}
	// 访问 User 字段
	user := LoginUserDTO.User
	// 根据邮箱和密码查询，判断用户是否存在
	userModel, err := SelectUser(&UserModel{Email: user.Email, Password: user.Password})
	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("用户登录", errors.New("邮箱或密码错误，请检查！")))
		return
	}
	// 转换为VO
	userVO := UserVO{
		Username: userModel.Username,
		Email:    userModel.Email,
		Bio:      userModel.Bio,
		Image:    userModel.Image,
		Token:    common.GenToken(userModel.ID),
	}

	c.JSON(http.StatusOK, gin.H{"user": userVO})
}

// 用户注册
func UsersRegist(c *gin.Context) {

	// 校验数据
	if err := c.ShouldBindJSON(&RegistUserDTO); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("用户注册", errors.New("请确认数据填写是否规范！")))
		return
	}
	// 取到 User
	var user = RegistUserDTO.User
	// 保存用户
	if err := SaveUser(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	// TODO：还需要取到完整数据
	userVO := UserVO{
		Username: user.Username,
		Email:    user.Email,
		Bio:      "",
		Image:    "",
		Token:    "",
	}
	c.JSON(http.StatusCreated, gin.H{"user": userVO})
}

// 关注用户
func FollowUser(c *gin.Context) {
}

// 取消关注用户
func UnFollowUser(c *gin.Context) {
}
