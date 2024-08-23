package model

// MenuItem 表示菜单项
type MenuItem struct {
	FamilyID int `gorm:"primaryKey" json:"family_id"` // 所属家庭ID
	FoodID   int `gorm:"primaryKey" json:"food_id"`   // 食品ID
	Quantity int `json:"quantity"`                    // 数量
}
