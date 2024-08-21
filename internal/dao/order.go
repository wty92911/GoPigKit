package dao

import (
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
)

// CreateOrder 创建新的订单或当前菜单
func CreateOrder(familyID int, items []model.OrderItem, status string) error {
	order := model.Order{
		FamilyID: familyID,
		Items:    items,
	}
	return database.DB.Create(&order).Error
}

// GetOrderByID 根据ID获取订单
func GetOrderByID(orderID uint) (*model.Order, error) {
	var order model.Order
	if err := database.DB.Preload("Items.Food").Order("created_at DESC").
		First(&order, orderID).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// GetOrdersByFamilyID 根据家庭ID获取所有订单
func GetOrdersByFamilyID(familyID int, status string) ([]model.Order, error) {
	var orders []model.Order
	if err := database.DB.Where("family_id = ? AND status = ?", familyID, status).Preload("Items.Food").
		Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// UpdateOrderStatus 更新订单状态
func UpdateOrderStatus(orderID uint, status string) error {
	return database.DB.Model(&model.Order{}).Where("id = ?", orderID).Update("status", status).Error
}
