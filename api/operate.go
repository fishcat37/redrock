package api

import (
	"redrock/dao"
	"redrock/model"
	"redrock/services"

	"github.com/gin-gonic/gin"
)

func Order(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_, err := services.ParseToken(token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	var order model.Order
	if err = c.BindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = dao.AddOrder(&order)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"info": "success", "status": 10000, "order_id": order.ID})
}
