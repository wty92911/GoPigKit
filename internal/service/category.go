package service

import (
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
)

// GetCategories 获取分类列表
func GetCategories(familyID uint) ([]model.Category, error) {
	return dao.GetCategories(familyID)
}

// CreateCategory 创建分类
/*
	给定对应的分类名称和图片，返回创建好的分类模型，使用transaction保证一致性
*/
func CreateCategory(category *model.Category) (*model.Category, error) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := dao.CreateCategory(tx, category); err != nil {
			return err // 返回错误时，事务会自动回滚
		}
		return nil // 返回 nil 时，事务会提交
	})

	if err != nil {
		return nil, err
	}
	return category, nil
}

// DeleteCategory 删除分类
/*
	给定category id，删除对应的内容，使用transaction保证一致性
*/
func DeleteCategory(id uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := dao.DeleteCategory(tx, id); err != nil {
			return err // 返回错误时，事务会自动回滚
		}
		return nil // 返回 nil 时，事务会提交
	})
}
