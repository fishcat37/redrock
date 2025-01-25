package dao

import (
	"redrock/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBUser *gorm.DB

func InitDB() error {
	dsn := "root:072231@tcp(127.0.0.1:3306)/first?charset=utf8&parseTime=True&loc=Local"
	var err error
	DBUser, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DBUser.AutoMigrate(&model.User{})
	return nil
}
