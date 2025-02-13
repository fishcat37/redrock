package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"redrock/config"
	"redrock/model"
)

var DB *gorm.DB

func InitSQL() error {
	var err error
	DB, err = gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&model.User{}, &model.Product{}, &model.Cart{}, &model.Comment{}, &model.Order{}, &model.OrderedProduct{})
	if err != nil {
		return err
	}
	return nil
}
