package service

import (
	"github.com/wty92911/GoPigKit/internal/dao"
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
)

// CreateOrder 创建订单
func CreateOrder(order *model.Order) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := dao.CreateOrder(tx, order); err != nil {
			return err
		}
		return nil
	})
}

// DeleteOrder 删除订单
func DeleteOrder(orderID uint) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := dao.DeleteOrder(tx, orderID); err != nil {
			return err
		}
		return nil
	})
}
