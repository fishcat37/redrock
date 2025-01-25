package api

import (
	"net/http"
	"redrock/dao"
	"redrock/model"
	"redrock/services"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	existingUser, isTrue := dao.FindUser(&user)
	if existingUser {
		c.JSON(400, gin.H{"error": "用户已存在"})
		return
	} else if !isTrue {
		c.JSON(500, gin.H{"error": "查找出错"})
		return
	}
	id, err := dao.AddUser(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"status":  "10000",
		"info":    "success",
		"message": "注册成功",
		"id":      id,
	})
}

func Token(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	info, _ := dao.FindUser(&user)
	if !info {
		c.JSON(400, gin.H{"error": "用户不存在或查找错误"})
		return
	}
	token, err := services.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成访问令牌失败"})
		return
	}
	refreshToken, err := services.CreateRefreshToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成刷新令牌失败"})
		return
	}
	c.JSON(200, gin.H{
		"status":        "10000",
		"info":          "success",
		"refresh_token": refreshToken,
		"token":         token,
	})

}

func Refresh(c *gin.Context) {
	var myToken model.Token
	myToken.Token = c.GetHeader("Authorization")
	if err := c.BindJSON(&myToken); err != nil {
		return
	}
	claims1, err1 := services.ParseToken(myToken.Token)
	claims2, err2 := services.ParseToken(myToken.RefreshToken)
	if err1 != nil || err2 != nil {
		c.JSON(401, gin.H{
			"message": "token无效",
		})
		return
	}
	if time.Now().Unix() > claims2.ExpiresAt.Time.Unix() {
		c.JSON(401, gin.H{
			"message": "refresh_token已过期",
		})
		return
	}
	info := model.User{
		Username: claims1.Username,
		Password: claims1.Password,
	}
	token, err := services.CreateToken(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成访问令牌失败"})
		return
	}
	refreshToken, err := services.CreateRefreshToken(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成刷新令牌失败"})
		return
	}
	c.JSON(200, gin.H{
		"status":        "10000",
		"info":          "success",
		"refresh_token": refreshToken,
		"token":         token,
	})
}

func Password(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := services.ParseToken(token)
	if err != nil {
		c.JSON(401, gin.H{
			"message": "token无效",
		})
		return
	}
	user := model.User{
		Username: claims.Username,
		Password: claims.Password,
	}
	flag, _ := dao.FindUser(&user)
	if !flag {
		c.JSON(400, gin.H{"error": "用户不存在或查找出错"})
		return
	}
	var password model.Password
	if err := c.BindJSON(&password); err != nil {
		return
	}
	if err := dao.UpdatePassword(&user, password); err != nil {
		c.JSON(500, gin.H{"error": "修改密码失败"})
		return
	}
	c.JSON(200, gin.H{
		"status":  "10000",
		"info":    "success",
		"message": "修改密码成功",
	})
}

func GetInfo(c *gin.Context) {
	var user model.User
	if err := c.BindUri(&user); err != nil {
		return
	}
	token := c.GetHeader("Authorization")
	claims, err := services.ParseToken(token)
	flag := dao.FindByID(&user)
	if !flag {
		c.JSON(500, gin.H{"message": "查找错误或用户不存在"})
		return
	}
	if err != nil || claims.Username != user.Username || claims.Password != user.Password {
		c.JSON(401, gin.H{"message": "Token 错误"})
		return
	}
	c.JSON(200, gin.H{"status": 10000, "info": "success", "data": user})
}

func UpdateInfo(c *gin.Context) {

}
