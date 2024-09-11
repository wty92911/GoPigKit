package dao

import (
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
)

// AddMenuItem 添加菜单项
func AddMenuItem(tx *gorm.DB, menuItem *model.MenuItem) error {
	return tx.Create(menuItem).Error
}

// UpdateMenuItem 更新菜单项数量
func UpdateMenuItem(tx *gorm.DB, menuItem *model.MenuItem) error {
	return tx.Save(menuItem).Error
}

// DeleteMenuItem 删除菜单项
func DeleteMenuItem(tx *gorm.DB, menuItemID uint) error {
	return tx.Delete(&model.MenuItem{}, menuItemID).Error
}
