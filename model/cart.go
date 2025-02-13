package model

type Cart struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	UserID    uint    `gorm:"user_id"`
	ProductID uint    `gorm:"product_id"`
	Quantity  uint    `gorm:"default:1"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}
