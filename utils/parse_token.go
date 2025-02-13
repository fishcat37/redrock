package utils

import (
	"fmt"
	"redrock/config"
	"redrock/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*model.CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}
	if claims.ExpiresAt.Unix() < time.Now().Unix() {
		return nil, fmt.Errorf("token is expired")
	}
	if claims.Issuer != config.Issuer || claims.Subject != claims.Username {
		return nil, fmt.Errorf("token is invalid")
	}
	return nil, err
}
