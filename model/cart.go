package model

import ()

type Cart struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	UserID    uint    `gorm:"user_id"`
	ProductID uint    `gorm:"product_id"`
	quantity  uint    `gorm:"default:1"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`
}
