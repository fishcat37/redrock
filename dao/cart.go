package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"redrock/model"
)

var DBCart *gorm.DB

func InitCart(dsn string) error {
	var err error
	DBCart, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = DBCart.AutoMigrate(&model.Cart{})
	if err != nil {
		return err
	}
	return nil
}

func AddCart(user model.User, product model.Product) error {
	var cart model.Cart
	result := DBCart.Model(&cart).Where("user_id = ? AND product_id = ?", user.ID, product.ID).First(&cart)
	if result.RowsAffected == 0 {
		DBCart.Model(&cart).Create(&model.Cart{UserID: user.ID, ProductID: product.ID})
	}
	result = DBCart.Model(&cart).Where("user_id = ? AND product_id = ?", user.ID, product.ID).Update("quantity", gorm.Expr("quantity + 1"))
	return result.Error
}
func GetCartProduct(user model.User) ([]model.Cart, error) {
	var cart []model.Cart
	result := DBCart.Model(&cart).Preload("Product").Where("user_id = ?", user.ID).Find(&cart)
	if result.RowsAffected == 0 {
		return cart, result.Error
	}
	return cart, nil
}
