package dao

import (
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
)

// GetMenuItems 获取菜单项，根据familyID
func GetMenuItems(familyID uint) ([]*model.MenuItem, error) {
	var menuItems []*model.MenuItem
	if err := database.DB.Where("family_id = ?", familyID).Find(&menuItems).Error; err != nil {
		return nil, err
	}
	return menuItems, nil
}

// AddMenuItem 添加菜单项
func AddMenuItem(tx *gorm.DB, menuItem *model.MenuItem) error {
	return tx.Create(menuItem).Error
}

// UpdateMenuItem 更新菜单项数量
func UpdateMenuItem(tx *gorm.DB, menuItem *model.MenuItem) error {
	return tx.Save(menuItem).Error
}

// DeleteMenuItem 删除菜单项
func DeleteMenuItem(tx *gorm.DB, familyID, foodID uint) error {
	return tx.Delete(&model.MenuItem{}, "family_id = ? AND food_id = ?", familyID, foodID).Error
}
