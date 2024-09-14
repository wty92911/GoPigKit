package model

// MenuItem 表示菜单项
type MenuItem struct {
	FamilyID  *uint   `gorm:"primaryKey;not null" json:"family_id"`                  // 所属家庭ID
	FoodID    *uint   `gorm:"primaryKey;not null" json:"food_id" binding:"required"` // 食品ID
	Quantity  uint    `json:"quantity" binding:"required"`                           // 数量
	CreatedBy *string `gorm:"type:varchar(255);not null" json:"created_by"`          // 创建者
	User      *User   `gorm:"foreignKey:CreatedBy;references:OpenID"`
}
