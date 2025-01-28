package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
	"redrock/model"
)

var DBUser *gorm.DB

func InitUsers(dsn string) error {
	var err error
	DBUser, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DBUser.AutoMigrate(&model.User{})
	return nil
}
func FindUser(user *model.User) (bool, bool) {
	result := DBUser.Model(user).Where("username = ?", user.Username).First(user)
	if result.Error == gorm.ErrRecordNotFound {

		return false, true
	} else if result.Error != nil {
		return false, false
	}
	return true, true
}
func AddUser(user model.User) (int, error) {
	result := DBUser.Model(&user).Create(&user)
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
	result := DBUser.Model(user).First(user, user.ID)
	if result.Error != nil {
		return false
	}
	return true
}
func UpdateUser(user *model.User) error {
	err := DBUser.Model(user).Where("username = ?", user.Username).Updates(*user).Error
	if err != nil {
		return err
	}
	return nil
}
