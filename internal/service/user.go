package service

import (
	"fmt"
	"github.com/wty92911/GoPigKit/internal/model"
)

// TODO:
func GetAllUsers() ([]model.User, error) {
	return nil, nil
}

func GetUser(id int) (*model.User, error) {
	return nil, fmt.Errorf("user id:%d not found", id)
}
