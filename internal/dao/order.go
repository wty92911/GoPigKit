package dao

import (
	"github.com/wty92911/GoPigKit/internal/model"
	"gorm.io/gorm"
)

// AddOrderItem 添加菜单项
func AddOrderItem(tx *gorm.DB, orderItem *model.OrderItem) error {
	return tx.Create(&orderItem).Error
}

// UpdateOrderItem 更新菜单项数量
func UpdateOrderItem(tx *gorm.DB, orderItem *model.OrderItem) error {
	return tx.Save(orderItem).Error

}

// DeleteOrderItem 删除菜单项
func DeleteOrderItem(tx *gorm.DB, OrderItemID uint) error {
	return tx.Delete(&model.OrderItem{}, OrderItemID).Error
}

// CreateOrder 创建订单
func CreateOrder(tx *gorm.DB, order *model.Order) error {
	return tx.Create(order).Error
}

// DeleteOrder 删除订单
func DeleteOrder(tx *gorm.DB, orderID uint) error {
	return tx.Delete(&model.Order{}, orderID).Error
}
