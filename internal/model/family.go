package model

import "gorm.io/gorm"

type Family struct {
	gorm.Model
	Name        string      `json:"name" binding:"required"`
	OwnerOpenID *string     `gorm:"type:varchar(255);uniqueIndex;not null" json:"owner_open_id"`
	Owner       *User       `gorm:"foreignKey:OwnerOpenID;references:OpenID" json:"owner"`
	Users       []*User     `gorm:"foreignKey:FamilyID" json:"users"`
	Orders      []*Order    `gorm:"foreignKey:FamilyID" json:"orders"`
	MenuItems   []*MenuItem `gorm:"foreignKey:FamilyID" json:"menu_items"`
	Categories  []*Category `gorm:"foreignKey:FamilyID" json:"categories"`
}

func (Family) TableName() string {
	return "families"
}
