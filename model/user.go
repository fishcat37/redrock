package model

import (
	"time"
)

type User struct {
	ID           uint      `json:"id" uri:"user_id" gorm:"primary_key"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Telephone    string    `json:"telephone"`
	QQ           string    `json:"qq"`
	Gender       string    `json:"gender"`
	Email        string    `json:"email"`
	Birthday     time.Time `json:"birthday" gorm:"type:date;default:'2000-01-01'"`
}
type Token struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
type Password struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
