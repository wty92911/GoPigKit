package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	TopName  string `gorm:"primaryKey" json:"top_name"`
	MidName  string `gorm:"primaryKey" json:"mid_name"`
	Name     string `gorm:"primaryKey" json:"name"`
	FamilyID uint   `json:"family_id"`
	Foods    []Food `gorm:"foreignKey:CategoryID" json:"foods"`
	ImageURL string `json:"image_url"`
}
