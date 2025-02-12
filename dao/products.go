package dao

import (
	"redrock/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "gorm.io/gorm"
)

var DBProduct *gorm.DB

func InitProduct(dsn string) error {
	var err error
	DBProduct, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = DBProduct.AutoMigrate(&model.Product{})
	if err != nil {
		return err
	}
	return nil
}

func GetProductList(products *[]model.Product) error {
	result := DBProduct.Model(&model.Product{}).Find(products)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindProduct(product *model.Product) error {
	result := DBProduct.Model(product).Where("name = ? OR id = ?", product.Name, product.ID).First(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindProductById(product *model.Product) error {
	result := DBProduct.Model(product).Where("id = ?", product.ID).First(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
