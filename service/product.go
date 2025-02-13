package service

import (
	"net/http"
	"redrock/dao"
	"redrock/model"
	"redrock/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) { //获取商品列表
	var products []model.Product
	if err := dao.GetProductList(&products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 10000, "info": "success", "data": products})

}

func Search(c *gin.Context) { //搜索商品
	token := c.GetHeader("AUthorization")
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token验证错误"})
	}
	user := model.User{ID: claims.ID, Username: claims.Username}
	existingUser, isTrue := dao.FindUser(&user)
	if !existingUser {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	} else if !isTrue {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查找出错"})
		return
	}
	var product model.Product
	var isGet bool
	product.Name, isGet = c.GetQuery("product_name")
	if !isGet {
		c.JSON(http.StatusBadRequest, gin.H{"error": "未传入商品"})
	}
	err = dao.FindProduct(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": 10000, "info": "success", "data": product})
}

func AddCart(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token解析出错"})
		return
	}
	user := model.User{
		ID:       claims.ID,
		Username: claims.Username,
	}
	var product model.Product
	var id uint64
	id, err = strconv.ParseUint(c.PostForm("product_id"), 10, 0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	product.ID = uint(id)
	err = dao.FindProduct(&product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	err = dao.AddCart(user, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 10000, "info": "success"})
}

func Cart(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := model.User{}
	// var id uint64

	// id, err = strconv.ParseUint(c.PostForm("user_id"), 10, 0)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// user.Username = claims.Username
	// user.ID = uint(id)
	err = c.BindQuery(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = claims.ID
	user.Username = claims.Username
	var carts []model.Cart
	err = dao.GetCartProduct(user, &carts)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 10000, "info": "success", "data": carts})
}

func GetProductInfo(c *gin.Context) {
	var product model.Product
	if err := c.BindUri(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := dao.FindProduct(&product); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 10000, "info": "success", "data": product})
}

func GetInfoByType(c *gin.Context) {
	var product model.Product
	if err := c.BindUri(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var products []model.Product
	if err := dao.FindProductByType(product, &products); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 10000, "info": "success", "data": products, "type": product.Type})
}
