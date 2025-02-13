package dao

import (
	"gorm.io/gorm"
	"redrock/model"
)

func AddOrder(order *model.Order) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		for i := range order.Products {
			order.Products[i].OrderID = order.ID
			order.Products[i].ID = 0
		}
		if err := tx.Create(&order.Products).Error; err != nil {
			return err
		}
		return nil
	})
}
