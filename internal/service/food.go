package service

import (
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/model"
)

// GetFoodsByCategoryID 根据分类ID获取菜品列表
func GetFoodsByCategoryID(categoryID uint) ([]*model.Food, error) {
	return dao.GetFoodsByCategoryID(categoryID)
}
