package model

import (
	"time"
)

type User struct {
	ID           int       `json:"id" gorm:"primary_key" uri:"user_id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Telephone    string    `json:"telephone"`
	QQ           string    `json:"qq"`
	Gender       string    `json:"gender"`
	Email        string    `json:"email"`
	Birthday     time.Time `json:"birthday" gorm:"type:date;default:CURRENT_DATE"`
}
type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
type Password struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
