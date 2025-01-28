package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"redrock/model"
	// "gorm.io/gorm"
)

var DBProduct *gorm.DB

func InitProduct(dsn string) error {
	var err error
	DBProduct, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DBProduct.AutoMigrate(&model.Product{})
	return nil
}

func GetProductList(products []model.Product) error {
	result := DBProduct.Model(&model.Product{}).Find(&products)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindProduct(product *model.Product) error {
	result := DBProduct.Model(product).Where("name = ?", product.Name).First(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
