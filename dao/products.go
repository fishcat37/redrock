package dao

import (
	"fmt"
	"redrock/model"
)

func GetProductList(products *[]model.Product) error {
	result := DB.Model(&model.Product{}).Find(products)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindProduct(product *model.Product) error {
	result := DB.Model(product).Where("name = ? OR id = ?", product.Name, product.ID).First(product)
	return result.Error
}

func FindProductById(id uint) error {
	product := &model.Product{}
	result := DB.Model(&model.Product{}).Where("id = ?", id).First(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindProductByType(product model.Product, products *[]model.Product) error {
	result := DB.Model(&model.Product{}).Where("type = ?", product.Type).Find(&products)
	if result.RowsAffected == 0 {
		return fmt.Errorf("未查到" + product.Type)
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
