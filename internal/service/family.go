package service

import (
	"fmt"
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
)

// GetAllFamilies 获得所有家庭
func GetAllFamilies() ([]*model.Family, error) {
	families, err := dao.GetAllFamilies()
	return families, err
}

// GetFamilyWithPreloads 获得家庭详细情况
func GetFamilyWithPreloads(id uint, preloads []string) (*model.Family, error) {
	family, err := dao.GetFamilyWithPreloads(id, preloads)
	return family, err
}

// CreateFamily 创建家庭
func CreateFamily(openID string, name string) (*model.Family, error) {
	user, err := dao.GetUser(openID)
	if err != nil {
		return nil, err
	}
	if user.FamilyID != nil {
		return nil, fmt.Errorf("user already in family %d", user.FamilyID)
	}

	tx := database.DB.Begin()
	family := &model.Family{
		Name:        name,
		OwnerOpenID: &openID,
	}
	err = dao.CreateFamily(tx, family)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	user.FamilyID = &family.ID
	err = dao.UpdateUser(tx, user)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return family, err
}

// JoinFamily 加入家庭
func JoinFamily(id uint, openID string) (*model.Family, error) {
	user, err := dao.GetUser(openID)
	if err != nil {
		return nil, err
	}
	if user.FamilyID != nil {
		return nil, fmt.Errorf("user already in family %d", user.FamilyID)
	}
	family, err := dao.GetFamily(id)
	if err != nil {
		return nil, err
	}

	tx := database.DB.Begin()
	user.FamilyID = &family.ID
	err = dao.UpdateUser(tx, user)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return family, err
}
