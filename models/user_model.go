package models

import (
	"my-realworld-go/common" // 本地的common包

	"github.com/jinzhu/gorm"
)

type UserModel struct {
	ID           uint    `gorm:"primary_key"`
	Username     string  `json:"username" gorm:"column:username"`
	Email        string  `json:"email" gorm:"column:email;unique_index"`
	Bio          string  `json:"bio" gorm:"column:bio;size:1024"`
	Image        *string `json:"image" gorm:"column:image"`
	PasswordHash string  `gorm:"column:password;not null"`
}

// 信息
type Profile struct {
	Username  string  `json:"username"`
	Bio       string  `json:"bio"`
	Image     *string `json:"image"`
	Following bool    `json:"following"`
}

// 数据表表名
func (UserModel) TableName() string {
	return "user"
}

type FollowModel struct {
	gorm.Model
	Following    UserModel
	FollowingID  uint
	FollowedBy   UserModel
	FollowedByID uint
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
	err := db.Save(data).Error
	return err
}
