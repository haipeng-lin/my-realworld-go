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
	// 转换为 ProfileVO
	profileVO := ProfileVO{
		Username: userModel.Username,
		Bio:      userModel.Bio,
		Image:    userModel.Image,
		// TODO：先默认为false，待改
		Following: false,
	}

	c.JSON(http.StatusOK, gin.H{"profile": profileVO})
}

// 用户登录
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
	// 修改用户上下文
	UpdateUserModelContext(c, userModel.ID)
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
	// 获取 User
	user := RegistUserDTO.User
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

// 获取用户
func GetUser(c *gin.Context) {
	// 从上下文获取用户
	currentUserModel := c.MustGet("current_user_model").(UserModel)
	// 返回用户VO
	userVO := UserVO{
		Username: currentUserModel.Username,
		Email:    currentUserModel.Email,
		Bio:      currentUserModel.Bio,
		Image:    currentUserModel.Image,
		Token:    common.GenToken(currentUserModel.ID),
	}
	c.JSON(http.StatusOK, gin.H{"user": userVO})
}

// 用户修改
func UserUpdate(c *gin.Context) {

	// 校验数据
	if err := c.ShouldBindJSON(&UpdateUserDTO); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("用户修改", errors.New("请确认数据填写是否规范！")))
		return
	}
	// 获取 user
	user := UpdateUserDTO.User
	// 从用户上下文获取用户ID
	currentUserModel := c.MustGet("current_user_model").(UserModel)
	user.ID = currentUserModel.ID
	// 修改用户
	if err := UpdateUser(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("数据库修改失败", err))
		return
	}
	// 修改用户上下文
	UpdateUserModelContext(c, currentUserModel.ID)

	// 根据用户ID查询完整信息
	fullUser, err := SelectUser(&UserModel{ID: user.ID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewError("获取用户信息失败", err))
		return
	}
	// 返回用户VO
	userVO := UserVO{
		Username: fullUser.Username,
		Email:    fullUser.Email,
		Bio:      fullUser.Bio,
		Image:    fullUser.Image,
		Token:    common.GenToken(user.ID),
	}
	c.JSON(http.StatusOK, gin.H{"user": userVO})
}

// 关注用户
func FollowUser(c *gin.Context) {
	// 解析请求参数 username
	username := c.Param("username")
	followedUserModel, err := SelectUser(&UserModel{Username: username})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("无效用户名！")))
		return
	}
	// 查询当前用户
	currentUserModel := c.MustGet("current_user_model").(UserModel)
	// 当前用户关注！
	err = currentUserModel.following(followedUserModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	// 返回被关注者个人资料
	profileVO := ProfileVO{
		Username:  followedUserModel.Username,
		Bio:       followedUserModel.Bio,
		Image:     followedUserModel.Image,
		Following: true,
	}
	c.JSON(http.StatusOK, gin.H{"profile": profileVO})
}

// 取消关注用户
func UnFollowUser(c *gin.Context) {
}
