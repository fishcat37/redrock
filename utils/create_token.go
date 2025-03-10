package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"redrock/config"
	"redrock/dao"
	"redrock/model"
	"time"
)

func CreateToken(user model.User) (string, error) {
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)), //1小时过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    config.Issuer, //签发者
			Subject:   user.Username, //主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtSecret)
}
