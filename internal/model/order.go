package model

import (
	"gorm.io/gorm"
)

// Order 表示一个订单
type Order struct {
	gorm.Model
	FamilyID uint        `gorm:"index" json:"family_id"`                                       // 家庭ID，外键
	Items    []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;" json:"items"` // 关联的订单项，联级删除

}

// OrderItem 表示订单中的一个项
type OrderItem struct {
	OrderID   uint `gorm:"primaryKey" json:"order_id"` // 订单ID，外键
	FoodID    int  `gorm:"primaryKey" json:"food_id"`  // 食品ID
	Quantity  int  `json:"quantity"`                   // 数量
	CreatedBy string
	User      User `gorm:"foreignKey:CreatedBy;references:OpenID"`
}
