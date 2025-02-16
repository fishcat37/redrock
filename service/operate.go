package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redrock/config"
	"redrock/dao"
	"redrock/model"
)

// Order handles the creation of a new order. It expects a JSON payload representing
// the order details in the request body and an "Authorization" token in the header.
// It performs the following steps:
// 1. Retrieves the "Authorization" token from the request header.
// 2. Parses the token to validate it.
// 3. Binds the JSON payload to an Order struct.
// 4. Adds the order to the database using the dao.AddOrder function.
// 5. Returns a JSON response indicating success or failure.
//
// If any step fails, it returns a JSON response with an appropriate info message
// and HTTP status code.
//
// Parameters:
// - c (*gin.Context): The Gin context for the request.
//
// Example response on success:
//
//	{
//	  "info": "success",
//	  "status": config.SuccessCode,
//	  "order_id": 123
//	}
//
// Example response on info:
//
//	{
//	  "info": "info message"
//	}
func Order(c *gin.Context) {
	//token := c.GetHeader("Authorization")
	//claims, err := utils.ParseToken(token)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{
	//		"status": config.TokenErrCode,
	//		"info":   err.Error()})
	//}
	id, get := c.Get("ID")
	if !get {
		c.JSON(http.StatusBadRequest, gin.H{"status": config.GetIDFailed, "error": "获取ID失败"})
		return
	}
	var order model.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   err.Error()})
		return
	}
	if order.UserID != id.(uint) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   "用户ID不匹配"})
		return
	}
	err := dao.AddOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": config.DataBaseErrCode,
			"info":   err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"info": "success", "status": config.SuccessCode, "data": gin.H{"order_id": order.ID}})
}

func GetOrder(c *gin.Context) {
	//token := c.GetHeader("Authorization")
	//claims, err := utils.ParseToken(token)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, gin.H{
	//		"status": config.TokenErrCode,
	//		"info":   err.Error()})
	//}
	id, get := c.Get("ID")
	if !get {
		c.JSON(http.StatusBadRequest, gin.H{"status": config.GetIDFailed, "info": "获取ID失败"})
		return
	}
	var order model.Order
	if err := c.ShouldBindUri(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   err.Error()})
		return
	}
	order.UserID = id.(uint)
	//if err := c.ShouldBindJSON(&order); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"status": config.RequestErrCode,
	//		"info":   err.Error()})
	//	return
	//}
	//if order.UserID != id {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"status": config.RequestErrCode,
	//		"info":   "用户ID不匹配"})
	//	return
	//}
	err := dao.GetOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": config.DataBaseErrCode,
			"info":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": config.SuccessCode, "data": order})
}

func GetOrderList(c *gin.Context) {
	id, get := c.Get("ID")
	if !get {
		c.JSON(http.StatusBadRequest, gin.H{"status": config.GetIDFailed, "error": "获取ID失败"})
		return
	}
	var order model.Order
	if err := c.ShouldBindUri(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   err.Error()})
		return
	}
	if id != order.UserID {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   "用户ID不匹配"})
		return
	}
	var orderList []model.Order
	err := dao.GetOrderList(&orderList, order.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": config.DataBaseErrCode,
			"info":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": config.SuccessCode, "data": orderList})
}

func UpdateOrder(c *gin.Context) {
	id, get := c.Get("ID")
	if !get {
		c.JSON(http.StatusBadRequest, gin.H{"status": config.GetIDFailed, "error": "获取ID失败"})
		return
	}
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   err.Error()})
		return
	}
	if order.UserID != id {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   "用户ID不匹配"})
		return
	}
	err := dao.UpdateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": config.DataBaseErrCode,
			"info":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": config.SuccessCode})
}

func DeleteOrder(c *gin.Context) {
	id, get := c.Get("ID")
	if !get {
		c.JSON(http.StatusBadRequest, gin.H{"status": config.GetIDFailed, "error": "获取ID失败"})
		return
	}
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   err.Error()})
		return
	}
	if order.UserID != id {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   "用户ID不匹配"})
		return
	}
	err := dao.DeleteOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": config.DataBaseErrCode,
			"info":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": config.SuccessCode})
}
