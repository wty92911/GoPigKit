package model

import (
	"gorm.io/gorm"
	"time"
)

// User 表示系统中的用户
type User struct {
	OpenID    string         `gorm:"primaryKey;type:varchar(255)" json:"open_id"`
	Name      string         `json:"name" validate:"required"`
	AvatarURL string         `json:"avatar_url"`
	FamilyID  uint           `json:"family_id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
