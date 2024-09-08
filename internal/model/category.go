package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	FamilyID *uint  `json:"family_id" gorm:"uniqueIndex:unique_combination;not null"`
	TopName  string `json:"top_name" gorm:"uniqueIndex:unique_combination, length:255"`
	MidName  string `json:"mid_name" gorm:"uniqueIndex:unique_combination, length:255"`
	Name     string `json:"name" gorm:"uniqueIndex:unique_combination, length:255"`
	ImageURL string `json:"image_url"`

	Foods []*Food `gorm:"foreignKey:CategoryID" json:"foods"`
}
