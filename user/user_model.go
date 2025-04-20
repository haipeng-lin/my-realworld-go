package user

import (
	"my-realworld-go/common" // 本地的common包
)

// 用户
type UserModel struct {
	ID       uint   `gorm:"primary_key"`
	Username string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email;unique_index"`
	Bio      string `json:"bio" gorm:"column:bio;size:1024"`
	Image    string `json:"image" gorm:"column:image"`
	Password string `gorm:"column:password;not null"`
}

// 用户关注
type FollowModel struct {
	UserID         uint `gorm:"column:A"` // 用户ID
	FollowedUserID uint `gorm:"column:B"` // 被关注用户ID
}

// 个人资料
type ProfileVO struct {
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

// 修改用户DTO
var UpdateUserDTO struct {
	User struct {
		ID       uint
		Email    string `json:"email" binding:"email"`
		Username string `json:"username" bingding:""`
		Password string `json:"password" bingding:""`
		Image    string `json:"image" binding:""`
		Bio      string `json:"bio" binding:""`
	} `json:"user"`
}

// 数据表表名
func (UserModel) TableName() string {
	return "user"
}

// 查找用户
func SelectUser(condition interface{}) (UserModel, error) {
	db := common.GetDB()
	var model UserModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

// 保存用户
func SaveUser(data interface{}) error {
	db := common.GetDB()
	err := db.Table("user").Create(data).Error
	return err
}

/**
 * 修改用户
 * 	Updates(data)：更新多个字段；Update(字段名，data)：更新指定字段
 *
 */
func UpdateUser(data interface{}) error {
	db := common.GetDB()
	err := db.Table("user").Updates(data).Error
	return err
}

// 用户关注：用户 currentUserModel 关注了 followedUserModel
func (currentUserModel UserModel) following(followedUserModel UserModel) error {
	db := common.GetDB()
	var follow FollowModel
	//
	err := db.Table("_userfollows").FirstOrCreate(&follow, &FollowModel{
		UserID:         currentUserModel.ID, // 当前用户
		FollowedUserID: followedUserModel.ID,
	}).Error
	return err
}

// 用户取消关注：用户 currentUserModel 取消关注 followedUserModel
func (currentUserModel UserModel) unFollowing(followedUserModel UserModel) error {
	db := common.GetDB()
	err := db.Table("_userfollows").Where(FollowModel{
		UserID:         currentUserModel.ID, // 当前用户
		FollowedUserID: followedUserModel.ID,
	}).Delete(FollowModel{}).Error
	return err
}
