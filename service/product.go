package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock/config"
	"redrock/dao"
	"redrock/model"
	"redrock/utils"
	"strconv"
)

func List(c *gin.Context) { //获取商品列表
	var products []model.Product
	if err := dao.GetProductList(&products); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": config.DataBaseErrCode,
			"info":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": config.SuccessCode, "info": "success", "data": products})

}

func Search(c *gin.Context) { //搜索商品
	token := c.GetHeader("Authorization")
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": config.TokenErrCode,
			"info":   "token验证错误"})
	}
	user := model.User{ID: claims.ID, Username: claims.Username}
	existingUser, isTrue := dao.FindUser(&user)
	if !existingUser {
		c.JSON(http.StatusNotFound, gin.H{
			"status": config.UserNotExistCode,
			"info":   "用户不存在"})
		return
	} else if !isTrue {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": config.DataBaseErrCode,
			"info":   "查找出错"})
		return
	}
	var product model.Product
	var isGet bool
	product.Name, isGet = c.GetQuery("product_name")
	if !isGet {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   "未传入商品"})
	}
	err = dao.FindProduct(&product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": config.ProductNotExistCode,
			"info":   err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": config.SuccessCode, "info": "success", "data": product})
}

func AddCart(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": config.TokenErrCode,
			"info":   "token解析出错"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"info": err.Error()})
		return
	}
	product.ID = uint(id)
	err = dao.FindProduct(&product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": config.ProductNotExistCode,
			"info":   err.Error()})
		return
	}
	err = dao.AddCart(user, product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": config.DataBaseErrCode,
			"info":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": config.SuccessCode, "info": "success"})
}

func Cart(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.TokenErrCode,
			"info":   err.Error()})
		return
	}
	user := model.User{}
	// var id uint64

	// id, err = strconv.ParseUint(c.PostForm("user_id"), 10, 0)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"info": err.Error()})
	// 	return
	// }
	// user.Username = claims.Username
	// user.ID = uint(id)
	err = c.BindQuery(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   err.Error()})
		return
	}
	user.ID = claims.ID
	user.Username = claims.Username
	var carts []model.Cart
	err = dao.GetCartProduct(user, &carts)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": config.CartNotExistCode,
			"info":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": config.SuccessCode, "info": "success", "data": carts})
}

func GetProductInfo(c *gin.Context) {
	var product model.Product
	if err := c.BindUri(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   err.Error()})
		return
	}
	if err := dao.FindProduct(&product); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": config.ProductNotExistCode,
			"info":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": config.SuccessCode, "info": "success", "data": product})
}

func GetInfoByType(c *gin.Context) {
	var product model.Product
	if err := c.BindQuery(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   err.Error()})
		return
	}
	var products []model.Product
	if err := dao.FindProductByType(product, &products); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": config.ProductNotExistCode,
			"info":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": config.SuccessCode, "info": "success", "data": gin.H{"data": products, "type": product.Type}})
}
