package model

import (
	"gorm.io/gorm"
	"mime/multipart"
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

type CreateCategoryRequest struct {
	FamilyID uint                  `form:"family_id" binding:"required"`
	TopName  string                `form:"top_name" binding:"required"`
	MidName  string                `form:"mid_name" binding:"required"`
	Name     string                `form:"name" binding:"required"`
	Image    *multipart.FileHeader `form:"image" binding:"required"`
}
