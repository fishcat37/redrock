package dao

import (
	"errors"
	"fmt"
	"redrock/model"

	"gorm.io/gorm"
)

func FindUser(user *model.User) (bool, bool) {
	result := DB.Model(user).Where("username = ?", user.Username).First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, true
	} else if result.Error != nil {
		return false, false
	}
	return true, true
}
func AddUser(user model.User) (int, error) {
	result := DB.Model(&user).Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(user.ID), nil
}

func UpdatePassword(user *model.User, password model.Password) error {
	result := DB.Model(user).Where("username = ? AND password = ?", user.Username, password.OldPassword).Update("password", password.NewPassword)
	if result.RowsAffected == 0 {
		return fmt.Errorf("don't have this user")
	}
	return result.Error
}
func FindByID(user *model.User) bool {
	result := DB.Model(user).First(user, user.ID)
	if result.Error != nil {
		return false
	}
	return true
}
func UpdateUser(user *model.User) error {
	result := DB.Model(user).Where("username = ?", user.Username).Updates(user)
	if result.RowsAffected == 0 {
		return fmt.Errorf("don't have this user")
	}
	err := DB.Model(&model.User{}).First(user, "id = ?", user.ID).Error
	return err
}
