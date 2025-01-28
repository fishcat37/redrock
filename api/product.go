package api

import (
	"github.com/gin-gonic/gin"
	"redrock/dao"
	"redrock/model"
	"redrock/services"
	"strconv"
)

func List(c *gin.Context) {
	var products []model.Product
	if err := dao.GetProductList(products); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": 10000, "info": "success", "data": products})

}

func Search(c *gin.Context) {
	token := c.GetHeader("AUthorization")
	claims, err := services.ParseToken(token)
	if err != nil {
		c.JSON(400, gin.H{"error": "token验证错误"})
	}
	user := model.User{ID: claims.ID, Username: claims.Username}
	existingUser, isTrue := dao.FindUser(&user)
	if !existingUser {
		c.JSON(400, gin.H{"error": "用户不存在"})
		return
	} else if !isTrue {
		c.JSON(500, gin.H{"error": "查找出错"})
		return
	}
	var product model.Product
	var isGet bool
	product.Name, isGet = c.GetQuery("product_name")
	if !isGet {
		c.JSON(400, gin.H{"error": "未传入商品"})
	}
	err = dao.FindProduct(&product)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"status": 10000, "info": "success", "data": product})
}

func AddCart(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := services.ParseToken(token)
	if err != nil {
		c.JSON(400, gin.H{"error": "token解析出错"})
	}
	user := model.User{
		ID:       claims.ID,
		Username: claims.Username,
	}
	var product model.Product
	var id uint64
	id, err = strconv.ParseUint(c.PostForm("product_id"), 10, 0)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	product.ID = uint(id)
	err = dao.FindProduct(&product)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	err = dao.AddCart(user, product)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"status": 10000, "info": "success"})
}
