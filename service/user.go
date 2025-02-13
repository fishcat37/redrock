package service

import (
	"net/http"
	"redrock/dao"
	"redrock/model"
	"redrock/utils"
	// "time"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	existingUser, isTrue := dao.FindUser(&user)
	if existingUser {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户已存在"})
		return
	} else if !isTrue {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查找出错"})
		return
	}
	id, err := dao.AddUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
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
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在或查找错误"})
		return
	}
	token, err := utils.CreateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成访问令牌失败"})
		return
	}
	refreshToken, err := utils.CreateRefreshToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成刷新令牌失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "10000",
		"info":   "success",
		"data": gin.H{
			"refresh_token": refreshToken,
			"token":         token,
		},
	})
}

func Refresh(c *gin.Context) {
	var myToken model.Token
	myToken.Token = c.GetHeader("Authorization")
	if err := c.BindJSON(&myToken); err != nil {
		return
	}
	_, err1 := utils.ParseToken(myToken.Token)
	claims2, err2 := utils.ParseToken(myToken.RefreshToken)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "token无效",
		})
		return
	}
	// if time.Now().Unix() > claims2.ExpiresAt.Time.Unix() {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "refresh_token已过期",
	// 	})
	// 	return
	// }
	info := model.User{
		ID:       claims2.ID,
		Username: claims2.Username,
	}
	token, err := utils.CreateToken(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成访问令牌失败"})
		return
	}
	refreshToken, err := utils.CreateRefreshToken(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成刷新令牌失败"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "10000",
		"info":   "success",
		"data": gin.H{
			"refresh_token": refreshToken,
			"token":         token,
		},
	})
}

func Password(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "token无效",
		})
		return
	}
	user := model.User{
		ID:       claims.ID,
		Username: claims.Username,
	}
	flag, _ := dao.FindUser(&user)
	if !flag {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在或查找出错"})
		return
	}
	var password model.Password
	if err := c.BindJSON(&password); err != nil {
		return
	}
	if err := dao.UpdatePassword(&user, password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "输入密码错误或修改密码失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
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
	claims, err := utils.ParseToken(token)
	flag := dao.FindByID(&user)
	if !flag {
		c.JSON(http.StatusNotFound, gin.H{"message": "查找错误或用户不存在"})
		return
	}
	if err != nil || claims.Username != user.Username {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token 错误"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 10000, "info": "success", "data": user})
}

func UpdateInfo(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误"})
		return
	}
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token 错误"})
		return
	}
	user.ID = claims.ID
	user.Username = claims.Username
	err = dao.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "更新信息失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 10000, "info": "success", "data": user})

}
