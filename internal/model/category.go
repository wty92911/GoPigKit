package model

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	FamilyID uint   `json:"family_id" gorm:"uniqueIndex:unique_combination"`
	TopName  string `json:"top_name" gorm:"uniqueIndex:unique_combination"`
	MidName  string `json:"mid_name" gorm:"uniqueIndex:unique_combination"`
	Name     string `json:"name" gorm:"uniqueIndex:unique_combination"`
	ImageURL string `json:"image_url"`

	Foods []Food `gorm:"foreignKey:CategoryID" json:"foods"`
}
