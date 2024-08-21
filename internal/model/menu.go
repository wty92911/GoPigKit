package model

import (
	"gorm.io/gorm"
	"time"
)

// MenuItem 表示菜单项
type MenuItem struct {
	gorm.Model
	FamilyID int       `gorm:"index" json:"family_id"` // 所属家庭ID
	FoodID   int       `json:"food_id"`                // 食品ID
	Food     Food      `json:"food"`                   // 关联的食品信息 查询时Preload
	Quantity int       `json:"quantity"`               // 数量
	AddedAt  time.Time `json:"added_at"`               // 添加时间
}
