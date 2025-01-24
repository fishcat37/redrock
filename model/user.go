package model

type User struct {
	ID           int    `json:"id" gorm:"primary_key"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Nickname     string `json:"nickname"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	Telephone    string `json:"telephone"`
	Qq           string `json:"qq"`
	Gender       string `json:"gender"`
	Email        string `json:"email"`
	Birthday     string `json:"birthday"`
}
