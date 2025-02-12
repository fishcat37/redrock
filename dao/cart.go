package dao

import (
	"redrock/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBCart *gorm.DB

func InitCart(dsn string) error {
	var err error
	DBCart, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = DBCart.AutoMigrate(&model.Cart{}, &model.Product{})
	if err != nil {
		return err
	}
	return nil
}

func AddCart(user model.User, product model.Product) error {
	var cart model.Cart
	result := DBCart.Model(&cart).Where("user_id = ? AND product_id = ?", user.ID, product.ID).First(&cart)
	if result.RowsAffected == 0 {
		result = DBCart.Model(&cart).Create(&model.Cart{UserID: user.ID, ProductID: product.ID})
		return result.Error
	}
	result = DBCart.Model(&cart).Where("user_id = ? AND product_id = ?", user.ID, product.ID).Update("quantity", gorm.Expr("quantity + 1"))
	return result.Error
}
func GetCartProduct(user model.User) ([]model.Cart, error) {
	var carts []model.Cart
	result := DBCart.Model(&model.Cart{}).Preload("Product").Where("user_id = ?", user.ID).Find(&carts)
	if result.RowsAffected == 0 {
		return carts, result.Error
	}
	return carts, nil
}
