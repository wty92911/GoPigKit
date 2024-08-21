package model

import (
	"gorm.io/gorm"
)

// Order 表示一个订单
type Order struct {
	gorm.Model
	FamilyID int         `gorm:"index" json:"family_id"`          // 家庭ID，外键
	Items    []OrderItem `gorm:"foreignKey:OrderID" json:"items"` // 关联的订单项
}

// OrderItem 表示订单中的一个项
type OrderItem struct {
	gorm.Model
	OrderID  int  `gorm:"index" json:"order_id"` // 订单ID，外键
	FoodID   int  `json:"food_id"`               // 食品ID
	Food     Food `json:"food"`                  // 关联的食品信息
	Quantity int  `json:"quantity"`              // 数量
}
