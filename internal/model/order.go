package model

import (
	"gorm.io/gorm"
)

// Order 表示一个订单
type Order struct {
	gorm.Model
	FamilyID *uint        `gorm:"index;not null" json:"family_id"`
	Items    []*OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;" json:"items" binding:"required"`
}

// OrderItem 表示订单中的一个项
type OrderItem struct {
	OrderID   *uint   `gorm:"primaryKey;not null" json:"order_id"`
	FoodID    *uint   `gorm:"primaryKey;not null" json:"food_id" binding:"required"`
	Quantity  uint    `json:"quantity" binding:"required"`
	CreatedBy *string `gorm:"type:varchar(255);not null" json:"created_by"`
	User      *User   `gorm:"foreignKey:CreatedBy;references:OpenID"`
}
