package dao

import (
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
)

// GetFood 根据ID获取食品
func GetFood(id uint) (*model.Food, error) {
	var food model.Food
	if err := database.DB.First(&food, id).Error; err != nil {
		return nil, err
	}
	return &food, nil
}

// CreateFood 创建新的食品
func CreateFood(food *model.Food) error {
	return database.DB.Create(food).Error
}

// UpdateFood 更新食品信息
func UpdateFood(food *model.Food) error {
	return database.DB.Save(food).Error
}

// DeleteFood 根据id删除食品
func DeleteFood(id int) error {
	return database.DB.Delete(&model.Food{}, id).Error
}

// ListFoods 获取所有食品
func ListFoods() ([]model.Food, error) {
	var foods []model.Food
	if err := database.DB.Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}

// FindFoodsByFamilyID 根据FamilyID获取食品
func FindFoodsByFamilyID(familyID int) ([]model.Food, error) {
	var foods []model.Food
	if err := database.DB.Where("family_id = ?", familyID).Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}
