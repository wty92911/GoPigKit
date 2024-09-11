package service

import (
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"log"
)

// GetFoodsByCategoryID 根据分类ID获取菜品列表
func GetFoodsByCategoryID(categoryID uint) ([]*model.Food, error) {
	return dao.GetFoodsByCategoryID(categoryID)
}

// CreateFood 创建菜品
func CreateFood(food *model.Food) (*model.Food, error) {
	tx := database.DB.Begin()

	err := dao.CreateFood(tx, food)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// 提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Printf("Transaction commit failed: %v", err)
		return nil, err
	}
	return food, nil
}
