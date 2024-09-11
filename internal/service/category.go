package service

import (
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
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
	tx := database.DB.Begin()

	err := dao.CreateCategory(tx, category)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return category, nil
}

// DeleteCategory 删除分类
/*
	给定category id，删除对应的内容，使用transaction保证一致性
*/
func DeleteCategory(id uint) error {
	tx := database.DB.Begin()

	if err := dao.DeleteCategory(tx, id); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
