package dao

import (
	// "gorm.io/gorm"
	// "gorm.io/driver/mysql"
	"redrock/model"
)

func Add_user(user model.User) (int, error) {
	var id int
	err := DBUser.Create(&user).Pluck("id", &id).Error
	if err != nil {
		return 0, err
	}
	return id, nil
}

// func SelectByPassword(user model.User,password string) error{

// }
