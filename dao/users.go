package dao

import (
	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
	"redrock/model"
)

func FindUser(user *model.User) (bool, bool) {
	result := DBUser.Where("username = ? AND password = ?", user.Username, user.Password).First(user)
	if result.Error == gorm.ErrRecordNotFound {

		return false, true
	} else if result.Error != nil {
		return false, false
	}
	return true, true
}
func AddUser(user model.User) (int, error) {
	result := DBUser.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(user.ID), nil
}

func UpdatePassword(user *model.User, password model.Password) error {
	result := DBUser.Model(user).Where("username = ? AND password = ?", user.Username, password.OldPassword).Update("password", password.NewPassword)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func FindByID(user *model.User) bool {
	result := DBUser.First(user, user.ID)
	if result.Error != nil {
		return false
	}
	return true
}
