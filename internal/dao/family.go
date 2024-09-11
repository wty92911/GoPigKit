package dao

import (
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
)

// GetFamily 根据ID获取Family
func GetFamily(id uint) (*model.Family, error) {
	var family model.Family
	if err := database.DB.First(&family, id).Error; err != nil {
		return nil, err
	}
	return &family, nil
}

// GetAllFamilies 获取所有Family
func GetAllFamilies() ([]*model.Family, error) {
	var families []*model.Family
	if err := database.DB.Find(&families).Error; err != nil {
		return nil, err
	}
	return families, nil
}

// GetFamilyWithPreloads 根据ID获取Family并预加载关联关系
func GetFamilyWithPreloads(id uint, preloads []string) (*model.Family, error) {
	var family model.Family
	db := database.DB
	// 预加载关联关系
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	if err := db.First(&family, id).Error; err != nil {
		return nil, err
	}
	return &family, nil
}

// CreateFamily 创建新的Family
func CreateFamily(tx *gorm.DB, family *model.Family) error {
	return tx.Create(family).Error
}

// DeleteFamily 根据ID删除Family
func DeleteFamily(tx *gorm.DB, id uint) error {
	return tx.Delete(&model.Family{}, id).Error
}

// UpdateFamily 更新Family信息
func UpdateFamily(tx *gorm.DB, family *model.Family) error {
	return tx.Save(family).Error
}
