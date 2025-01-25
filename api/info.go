package api

// import (
// 	"redrock/dao"
// 	"redrock/model"
// 	"redrock/services"

// 	"github.com/gin-gonic/gin"
// )

// func GetInfo(c *gin.Context) {
// 	var user model.User
// 	if err := c.BindUri(&user); err != nil {
// 		return
// 	}
// 	token := c.GetHeader("Authorization")
// 	claims, err := services.ParseToken(token)
// 	flag := dao.FindByID(&user)
// 	if !flag {
// 		c.JSON(500, gin.H{"message": "查找错误或用户不存在"})
// 		return
// 	}
// 	if err != nil || claims.Username != user.Username || claims.Password != user.Password {
// 		c.JSON(401, gin.H{"message": "Token 错误"})
// 		return
// 	}
// 	c.JSON(200, gin.H{"status": 10000, "info": "success", "data": user})
// }
