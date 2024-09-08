package model

import (
	"gorm.io/gorm"
)

// Order 表示一个订单
type Order struct {
	gorm.Model
	FamilyID *uint        `gorm:"index;not null" json:"family_id"`                              // 家庭ID，外键
	Items    []*OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;" json:"items"` // 关联的订单项，联级删除

}

// OrderItem 表示订单中的一个项
type OrderItem struct {
	OrderID   *uint   `gorm:"primaryKey;not null" json:"order_id"`          // 订单ID，外键
	FoodID    *uint   `gorm:"primaryKey;not null" json:"food_id"`           // 食品ID
	Quantity  int     `json:"quantity"`                                     // 数量
	CreatedBy *string `gorm:"type:varchar(255);not null" json:"created_by"` // 创建者，外键
	User      *User   `gorm:"foreignKey:CreatedBy;references:OpenID"`
}
