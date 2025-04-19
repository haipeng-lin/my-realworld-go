package user

import (
	"my-realworld-go/common" // 本地的common包
)

// 用户
type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email;unique_index"`
	Bio      string `json:"bio" gorm:"column:bio;size:1024"`
	Image    string `json:"image" gorm:"column:image"`
	Password string `gorm:"column:password;not null"`
}

// 个人资料
type Profile struct {
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

// 用户VO
type UserVO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Token    string `json:"token"`
}

// 登录用户DTO
var LoginUserDTO struct {
	User struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	} `json:"user"`
}

// 注册用户DTO
var RegistUserDTO struct {
	User struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		Username string `json:"username" binding:"required"`
	} `json:"user"`
}

// 数据表表名
func (User) TableName() string {
	return "user"
}

// 查找用户
func SelectUser(condition interface{}) (User, error) {
	db := common.GetDB()
	var model User
	err := db.Where(condition).First(&model).Error
	return model, err
}

// 保存用户
func SaveUser(data interface{}) error {
	db := common.GetDB()
	err := db.Table("user").Create(data).Error
	return err
}
