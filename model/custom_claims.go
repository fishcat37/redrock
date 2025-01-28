package model

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	ID       uint   `json:"ID"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
