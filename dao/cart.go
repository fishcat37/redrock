package dao

import (
	"fmt"
	"redrock/model"

	"gorm.io/gorm"
)

func AddCart(user model.User, product model.Product) error {
	var cart model.Cart
	err := FindProduct(&product)
	if err != nil {
		return fmt.Errorf("product not found")
	}
	result := DB.Model(&cart).Where("user_id = ? AND product_id = ?", user.ID, product.ID).First(&cart)
	if result.RowsAffected == 0 {
		result = DB.Model(&cart).Create(&model.Cart{UserID: user.ID, ProductID: product.ID})
		return result.Error
	}
	result = DB.Model(&cart).Where("user_id = ? AND product_id = ?", user.ID, product.ID).Update("quantity", gorm.Expr("quantity + 1"))
	return result.Error
}
func GetCartProduct(user model.User, carts *[]model.Cart) error {
	result := DB.Model(&model.Cart{}).Preload("Product").Where("user_id = ?", user.ID).Find(&carts)
	if result.RowsAffected == 0 {
		return fmt.Errorf("cart product not found")
	}
	return result.Error
}
