package service

import (
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
)

// GetMenuItems 获取菜单列表，根据family_id
func GetMenuItems(familyID uint) ([]*model.MenuItem, error) {
	return dao.GetMenuItems(familyID)
}

// AddMenuItem 添加菜单项
func AddMenuItem(item *model.MenuItem) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := dao.AddMenuItem(tx, item); err != nil {
			return err // 返回错误时，事务会自动回滚
		}
		return nil // 返回 nil 时，事务会提交
	})
}

// UpdateMenuItem 更新菜单项
func UpdateMenuItem(menuItem *model.MenuItem) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := dao.UpdateMenuItem(tx, menuItem); err != nil {
			return err
		}
		return nil
	})
}

// DeleteMenuItem 删除菜单项
func DeleteMenuItem(familyID, foodID uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := dao.DeleteMenuItem(tx, familyID, foodID); err != nil {
			return err
		}
		return nil
	})
}
