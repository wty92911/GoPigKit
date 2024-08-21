package model

import (
	"gorm.io/gorm"
	"time"
)

// User 表示系统中的用户
type User struct {
	OpenID    string         `gorm:"primaryKey" json:"open_id"`
	Name      string         `json:"name" validate:"required"`
	FamilyID  int            `json:"family_id"`
	Role      string         `json:"role" validate:"oneof=owner member"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
