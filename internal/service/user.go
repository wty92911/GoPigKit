package service

import (
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/model"
)

func UpdateUser(user *model.User) error {
	return dao.UpdateUser(user)
}

func CreateUser(user *model.User) error {
	return dao.CreateUser(user)
}

func GetUser(openID string) (*model.User, error) {
	return dao.GetUser(openID)
}

func ExistUser(openID string) bool {
	_, err := dao.GetUser(openID)
	return err == nil
}
