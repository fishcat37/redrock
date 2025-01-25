package services

import (
	"redrock/config"
	"redrock/model"

	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
