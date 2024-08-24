package model

import "gorm.io/gorm"

type Family struct {
	gorm.Model
	Name        string     `json:"name" validate:"required"`
	OwnerOpenID string     `gorm:"uniqueIndex" json:"owner_open_id"`
	Users       []User     `gorm:"foreignKey:FamilyID" json:"users"`
	Foods       []Food     `gorm:"foreignKey:FamilyID" json:"foods"`
	Orders      []Order    `gorm:"foreignKey:FamilyID" json:"orders"`
	MenuItems   []MenuItem `gorm:"foreignKey:FamilyID" json:"menu_items"`
}

func (Family) TableName() string {
	return "families"
}
