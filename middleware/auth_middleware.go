package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"redrock/config"
	"redrock/model"
	"time"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JwtSecretKey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": config.TokenErrCode,
				"info":   "Invalid token"})
			c.Abort()
			return
		}
		claims, ok := token.Claims.(*model.CustomClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": config.TokenErrCode,
				"info":   "Invalid token"})
			c.Abort()
			return
		}
		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": config.TokenErrCode,
				"info":   "Token expired"})
			c.Abort()
			return
		} else if claims.Issuer != config.Issuer || claims.Subject != claims.Username {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": config.TokenErrCode,
				"info":   "Invalid token"})
			c.Abort()
			return
		}
		c.Set("ID", claims.ID)
		c.Set("Username", claims.Username)
		c.Next()
		return
	}
}
