package dao

import (
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
)

// GetFood 根据ID获取食品
func GetFood(id uint) (*model.Food, error) {
	var food model.Food
	if err := database.DB.First(&food, id).Error; err != nil {
		return nil, err
	}
	return &food, nil
}

// GetFoodsByFamilyID 根据FamilyID获取食品
func GetFoodsByFamilyID(familyID uint) ([]*model.Food, error) {
	var foods []*model.Food
	if err := database.DB.Where("family_id = ?", familyID).Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}

// GetFoodsByCategoryID 根据CategoryID获取食品
func GetFoodsByCategoryID(categoryID uint) ([]*model.Food, error) {
	var foods []*model.Food
	err := database.DB.Where("category_id = ?", categoryID).Find(&foods).Error
	return foods, err
}

// CreateFood 创建新的食品
func CreateFood(tx *gorm.DB, food *model.Food) error {
	return tx.Create(food).Error
}

// UpdateFood 更新食品信息
func UpdateFood(tx *gorm.DB, food *model.Food) error {
	return tx.Save(food).Error
}

// DeleteFood 根据id删除食品
func DeleteFood(tx *gorm.DB, id uint) error {
	return tx.Delete(&model.Food{}, id).Error
}
