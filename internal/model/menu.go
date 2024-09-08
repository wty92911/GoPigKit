package model

// MenuItem 表示菜单项
type MenuItem struct {
	FamilyID  *uint `gorm:"primaryKey;not null" json:"family_id"` // 所属家庭ID
	FoodID    *int  `gorm:"primaryKey;not null" json:"food_id"`   // 食品ID
	Quantity  int   `json:"quantity"`                             // 数量
	CreatedBy *string
	User      *User `gorm:"foreignKey:CreatedBy;references:OpenID"`
}
