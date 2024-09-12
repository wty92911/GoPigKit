package service

import (
	"fmt"
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
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
		return nil, fmt.Errorf("user already in family %d", *user.FamilyID)
	}

	var family *model.Family
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		family = &model.Family{
			Name:        name,
			OwnerOpenID: &openID,
		}

		if err := dao.CreateFamily(tx, family); err != nil {
			return err // 返回错误时，事务会自动回滚
		}

		user.FamilyID = &family.ID
		if err := dao.UpdateUser(tx, user); err != nil {
			return err // 返回错误时，事务会自动回滚
		}

		return nil // 返回 nil 时，事务会提交
	})

	if err != nil {
		return nil, err
	}
	return family, nil
}

// JoinFamily 加入家庭
func JoinFamily(id uint, openID string) (*model.Family, error) {
	user, err := dao.GetUser(openID)
	if err != nil {
		return nil, err
	}
	if user.FamilyID != nil {
		return nil, fmt.Errorf("user already in family %d", *user.FamilyID)
	}
	family, err := dao.GetFamily(id)
	if err != nil {
		return nil, err
	}

	err = database.DB.Transaction(func(tx *gorm.DB) error {
		user.FamilyID = &family.ID
		if err := dao.UpdateUser(tx, user); err != nil {
			return err // 返回错误时，事务会自动回滚
		}
		return nil // 返回 nil 时，事务会提交
	})

	if err != nil {
		return nil, err
	}
	return family, nil
}
