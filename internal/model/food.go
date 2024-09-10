package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

// Food 表示食品的结构体
type Food struct {
	gorm.Model
	Title       string          `gorm:"type:varchar(255);not null" json:"title" binding:"required"`      // 食品名称
	Price       int             `json:"price" binding:"required,gt=0"`                                   // 食品价格
	Desc        string          `gorm:"type:varchar(255);not null" json:"desc"`                          // 食品描述
	ImageURLs   json.RawMessage `gorm:"type:text" json:"image_urls" binding:"required"`                  // 食品图片
	CategoryID  *uint           `gorm:"index;not null" json:"category_id" binding:"required"`            // 食品分类ID，外键
	CreatedBy   *string         `gorm:"type:varchar(255);not null" json:"created_by" binding:"required"` // 创建者
	CreatedUser *User           `gorm:"foreignKey:CreatedBy;references:OpenID"`
}
