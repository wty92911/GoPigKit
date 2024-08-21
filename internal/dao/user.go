package dao

import (
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
)

// GetUser 根据 OpenID 获取用户
func GetUser(openID string) (*model.User, error) {
	var user model.User
	if err := database.DB.Where("open_id = ?", openID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser 创建新用户
func CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}

// DeleteUser 通过openID删除用户
func DeleteUser(openID string) error {
	return database.DB.Delete(&model.User{}, openID).Error
}

// UpdateUser 更新用户信息
func UpdateUser(user *model.User) error {
	return database.DB.Save(user).Error
}
