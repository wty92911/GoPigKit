package service

import (
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/model"
	"mime/multipart"
)

func GetCategories(familyID uint) ([]model.Category, error) {
	categories, err := dao.GetCategories(familyID)
	return categories, err
}

func CreateCategory(familyID uint, topName, midName, name string, file *multipart.FileHeader) (*model.Category, error) {
	return nil, nil
}
