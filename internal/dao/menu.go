package dao

import (
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
)

// AddMenuItem 添加菜单项
func AddMenuItem(familyID int, foodID int, quantity int) error {
	menuItem := model.MenuItem{
		FamilyID: familyID,
		FoodID:   foodID,
		Quantity: quantity,
	}
	return database.DB.Create(&menuItem).Error
}

// GetMenuItems 获取家庭的所有菜单项
func GetMenuItems(familyID int) ([]model.MenuItem, error) {
	var menuItems []model.MenuItem
	if err := database.DB.Where("family_id = ?", familyID).Preload("Food").Find(&menuItems).Error; err != nil {
		return nil, err
	}
	return menuItems, nil
}

// UpdateMenuItem 更新菜单项数量
func UpdateMenuItem(menuItemID uint, quantity int) error {
	return database.DB.Model(&model.MenuItem{}).Where("id = ?", menuItemID).Update("quantity", quantity).Error
}

// DeleteMenuItem 删除菜单项
func DeleteMenuItem(menuItemID uint) error {
	return database.DB.Delete(&model.MenuItem{}, menuItemID).Error
}
