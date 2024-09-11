package dao

import (
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
)

// GetCategory 根据id获取分类
func GetCategory(id uint) (*model.Category, error) {
	var category model.Category
	err := database.DB.First(&category, id).Error
	return &category, err
}

// GetCategories 根据familyID获取分类列表，按照ID升序
func GetCategories(familyID uint) ([]model.Category, error) {
	var categories []model.Category
	if err := database.DB.Where("family_id = ?", familyID).Order("id asc").Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

// GetCategoryWithPreloads 根据id获取分类，预加载
func GetCategoryWithPreloads(id uint, preloads []string) (*model.Category, error) {
	var category model.Category
	db := database.DB
	// 预加载关联关系
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	if err := db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

// CreateCategory 创建分类
func CreateCategory(tx *gorm.DB, category *model.Category) error {
	return tx.Create(category).Error
}

// UpdateCategory 更新分类
func UpdateCategory(tx *gorm.DB, category *model.Category) error {
	return tx.Save(category).Error
}

// DeleteCategory 根据id删除分类
func DeleteCategory(tx *gorm.DB, id uint) error {
	return tx.Delete(&model.Category{}, id).Error
}
