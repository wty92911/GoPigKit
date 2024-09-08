package model

import "gorm.io/gorm"

// Food 表示食品的结构体
type Food struct {
	gorm.Model
	Title     string `json:"title"`                 // 食品名称
	Price     int    `json:"price" validate:"gt=0"` // 食品价格
	Desc      string `json:"desc"`                  // 食品描述
	ImageURLs string `json:"image_urls"`            // 食品图片 json list

	CategoryID  *uint   `gorm:"index;not null" json:"category_id"`            // 食品分类ID，外键
	CreatedBy   *string `gorm:"type:varchar(255);not null" json:"created_by"` // 创建者
	CreatedUser *User   `gorm:"foreignKey:CreatedBy;references:OpenID"`
}
