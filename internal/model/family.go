package model

import "gorm.io/gorm"

type Family struct {
	gorm.Model
	Name    string `json:"name" validate:"required"`
	OwnerID int    `gorm:"uniqueIndex" json:"owner_id"`
	Users   []User `gorm:"foreignKey:FamilyID" json:"users"`
	Foods   []Food `gorm:"foreignKey:FamilyID" json:"foods"`
}

func (Family) TableName() string {
	return "families"
}
