package utils

import (
	"redrock/config"
	"redrock/dao"
	"redrock/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateRefreshToken(user model.User) (string, error) {
	existingUser, isTrue := dao.FindUser(&user)
	if !existingUser {
		panic("用户不存在")
	} else if !isTrue {
		panic("查找错误")
	}
	claims := model.CustomClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)), // 24 小时后过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    config.Issuer,
			Subject:   user.Username,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtSecret)
}
