package service

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
)

const categoryImagePrefix = "images/category"

// GetCategories 获取分类列表
func GetCategories(familyID uint) ([]model.Category, error) {
	return dao.GetCategories(familyID)
}

// CreateCategory 创建分类
/*
	给定对应的分类名称和图片，返回创建好的分类模型，使用transaction保证一致性
*/
func CreateCategory(category *model.Category) (*model.Category, error) {
	// 开启事务
	tx := database.DB.Begin()

	err := dao.CreateCategory(tx, category)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// 提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return category, nil
}

// DeleteCategory 删除分类
/*
	给定category id，删除对应的内容和图片对象，使用transaction保证一致性
*/
func DeleteCategory(id uint) error {
	category, err := dao.GetCategory(id)
	if err != nil {
		return err
	}
	// 开启事务
	tx := database.DB.Begin()

	// 删除数据库上的分类
	err = dao.DeleteCategory(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 删除MinIO上的图片
	err = database.MinIOClient.RemoveObject(
		context.Background(),
		database.MinIOBucket,
		category.ImageURL,
		minio.RemoveObjectOptions{ForceDelete: true},
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		// 这里不需要补偿事务，数据库的导致的错误返回之后，用户进行重复删除，重复删除时MinIO不会报错。
		return err
	}
	return nil
}
