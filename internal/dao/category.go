package dao

import (
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
)

// CreateCategory 创建分类
func CreateCategory(category *model.Category) error {
	return database.DB.Create(category).Error
}

// UpdateCategory 更新分类
func UpdateCategory(category *model.Category) error {
	return database.DB.Save(category).Error
}

// DeleteCategory 根据id删除分类
func DeleteCategory(id uint) error {
	return database.DB.Delete(&model.Category{}, id).Error
}

// GetCategories 根据familyID获取分类列表，按照ID升序
func GetCategories(familyID uint) ([]model.Category, error) {
	var categories []model.Category
	if err := database.DB.Where("family_id = ?", familyID).Order("id asc").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
