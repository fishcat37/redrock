package dao

import (
	"fmt"
	"gorm.io/gorm"
	"redrock/model"
)

func AddOrder(order *model.Order) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(order).Error; err != nil {
			return err
		}
		return nil
	})
}

func GetOrder(order *model.Order) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		result := DB.Model(order).Where("id = ? AND user_id = ?", order.ID, order.UserID).First(order)
		if result.Error != nil {
			return result.Error
		}
		result = DB.Model(&model.OrderedProduct{}).Where("order_id = ?", order.ID).Find(&order.Products)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func GetOrderList(orders *[]model.Order, userID uint) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		result := DB.Model(&model.Order{}).Where("user_id = ?", userID).Find(orders)
		if result.Error != nil {
			return result.Error
		}
		for i := range *orders {
			result := DB.Model(&model.OrderedProduct{}).Where("order_id = ?", (*orders)[i].ID).Find(&(*orders)[i].Products)
			if result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
}

func UpdateOrder(order model.Order) error {
	return DB.Transaction(func(tx *gorm.DB) error {

		if result := tx.Model(&order).Select("address").Where("id = ?", order.ID).Updates(order); result.Error != nil {
			if result.RowsAffected == 0 {
				return fmt.Errorf("未找到")
			}
			return result.Error
		}
		return nil
	})
}

func DeleteOrder(order *model.Order) error {
	err := GetOrder(order)
	if err != nil {
		return err
	}
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", order.ID).Preload("Products").Delete(&order).Error; err != nil {
			return err
		}
		return nil
	})
}
