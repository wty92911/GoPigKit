package services

import (
	"fmt"
	"github.com/wty92911/GoPigKit/internal/models"
)

// TODO:
func GetAllUsers() ([]models.User, error) {
	return nil, nil
}

func GetUser(id int) (*models.User, error) {
	return nil, fmt.Errorf("user id:%d not found", id)
}
