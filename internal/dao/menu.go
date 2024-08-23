package dao

import (
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
)

// AddMenuItem 添加菜单项
func AddMenuItem(menuItem *model.MenuItem) error {
	return database.DB.Create(menuItem).Error
}

// UpdateMenuItem 更新菜单项数量
func UpdateMenuItem(menuItem *model.MenuItem) error {
	return database.DB.Save(menuItem).Error
}

// DeleteMenuItem 删除菜单项
func DeleteMenuItem(menuItemID uint) error {
	return database.DB.Delete(&model.MenuItem{}, menuItemID).Error
}
