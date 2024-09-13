package service

import (
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
)

// GetAllFoods 获取指定家庭的所有菜品
func GetAllFoods(familyID uint) ([]*model.Food, error) {
	return dao.GetFoodsByFamilyID(familyID)
}

// GetFoodsByCategoryID 根据分类ID获取菜品列表
func GetFoodsByCategoryID(categoryID uint) ([]*model.Food, error) {
	return dao.GetFoodsByCategoryID(categoryID)
}

// CreateFood 创建菜品
func CreateFood(food *model.Food) (*model.Food, error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := dao.CreateFood(tx, food); err != nil {
			return err // 返回错误时，事务会自动回滚
		}
		return nil // 返回 nil 时，事务会提交
	})

	if err != nil {
		return nil, err
	}
	return food, nil
}

// DeleteFood 删除菜品
func DeleteFood(id uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := dao.DeleteFood(tx, id); err != nil {
			return err
		}
		return nil
	})
}
