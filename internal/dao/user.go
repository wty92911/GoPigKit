package dao

import (
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
)

// GetUser 根据 OpenID 获取用户
func GetUser(openID string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("open_id = ?", openID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUsers 获取用户列表
func GetUsers() ([]*model.User, error) {
	var users []*model.User
	if err := database.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// CreateUser 创建新用户
func CreateUser(tx *gorm.DB, user *model.User) error {
	return tx.Create(user).Error
}

// DeleteUser 通过openID删除用户
func DeleteUser(tx *gorm.DB, openID string) error {
	return tx.Delete(&model.User{OpenID: openID}).Error
}

// UpdateUser 更新用户信息
func UpdateUser(tx *gorm.DB, user *model.User) error {
	return tx.Save(user).Error
}
