package model

import "gorm.io/gorm"

// Food 表示食品的结构体
type Food struct {
	gorm.Model
	Title     string `json:"title"`                  // 食品名称
	Price     int    `json:"price" validate:"gt=0"`  // 食品价格
	Desc      string `json:"desc"`                   // 食品描述
	Images    string `json:"images"`                 // 食品图片
	FamilyID  uint   `gorm:"index" json:"family_id"` // 家庭ID，外键
	CreatedBy string
	User      User `gorm:"foreignKey:CreatedBy;references:OpenID"`
}
