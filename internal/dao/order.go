package dao

import (
	"github.com/wty92911/GoPigKit/internal/database"
	"github.com/wty92911/GoPigKit/internal/model"
)

// AddOrderItem 添加菜单项
func AddOrderItem(orderItem *model.OrderItem) error {
	return database.DB.Create(&orderItem).Error
}

// UpdateOrderItem 更新菜单项数量
func UpdateOrderItem(orderItem *model.OrderItem) error {
	return database.DB.Save(orderItem).Error

}

// DeleteOrderItem 删除菜单项
func DeleteOrderItem(OrderItemID uint) error {
	return database.DB.Delete(&model.OrderItem{}, OrderItemID).Error
}

// CreateOrder 创建订单
func CreateOrder(order *model.Order) error {
	return database.DB.Create(order).Error
}

// DeleteOrder 删除订单
func DeleteOrder(orderID uint) error {
	return database.DB.Delete(&model.Order{}, orderID).Error
}
