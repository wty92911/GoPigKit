package model

import (
	"gorm.io/gorm"
	"mime/multipart"
)

type Category struct {
	gorm.Model
	FamilyID *uint  `json:"family_id" gorm:"uniqueIndex:unique_combination;not null"`
	TopName  string `json:"top_name" gorm:"uniqueIndex:unique_combination, length:255" binding:"required"`
	MidName  string `json:"mid_name" gorm:"uniqueIndex:unique_combination, length:255" binding:"required"`
	Name     string `json:"name" gorm:"uniqueIndex:unique_combination, length:255" binding:"required"`
	ImageURL string `json:"image_url" gorm:"type:varchar(255)" binding:"required"`

	Foods []*Food `gorm:"foreignKey:CategoryID" json:"foods"`
}

type CreateCategoryRequest struct {
	FamilyID uint                  `form:"family_id" binding:"required"`
	TopName  string                `form:"top_name" binding:"required"`
	MidName  string                `form:"mid_name" binding:"required"`
	Name     string                `form:"name" binding:"required"`
	ImageURL *multipart.FileHeader `form:"image" binding:"required"`
}
