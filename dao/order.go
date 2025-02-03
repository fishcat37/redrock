package dao

import (
	"redrock/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBOrder *gorm.DB

func InitOrder(dsn string) error {
	var err error
	DBOrder, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = DBOrder.AutoMigrate(&model.Order{})
	if err != nil {
		return err
	}
	return nil
}

func AddOrder(order *model.Order) error {
	result := DBOrder.Model(&model.Order{}).Create(order)
	return result.Error
}
