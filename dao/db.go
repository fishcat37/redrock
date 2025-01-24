package dao

import (
	"redrock/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBUser *gorm.DB

func InitDB() {
	dsn := "root:072231@tcp(127.0.0.1:3306)/first?charset=utf8&parseTime=True&loc=Local"
	DBUser, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DBUser.AutoMigrate(&model.User{})
}
