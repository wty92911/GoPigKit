package service

import (
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
)

// GetUser 获取用户信息
func GetUser(openID string) (*model.User, error) {
	return dao.GetUser(openID)
}

// GetUsers 获取用户列表
func GetUsers() ([]*model.User, error) {
	return dao.GetUsers()
}

// CreateUser 创建用户
func CreateUser(user *model.User) error {
	tx := database.DB.Begin()
	if err := dao.CreateUser(tx, user); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// UpdateUser 更新用户信息
func UpdateUser(user *model.User) error {
	tx := database.DB.Begin()
	if err := dao.UpdateUser(tx, user); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

// ExistUser 判断用户是否存在
func ExistUser(openID string) bool {
	_, err := dao.GetUser(openID)
	return err == nil
}
