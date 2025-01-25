package services

import (
	"redrock/config"
	"redrock/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(user model.User) (string, error) {
	claims := model.CustomClaims{
		Username: user.Username,
		Password: user.Password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), //1小时过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    config.Issuer,
			Subject:   config.Subject,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtSecret)
}
