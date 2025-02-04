package api

import (
	"github.com/gin-gonic/gin"
	"redrock/dao"
	"redrock/model"
	"redrock/services"
)

func GetComment(c *gin.Context) {
	var comment model.Comment
	if err := c.BindUri(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	comments, err := dao.GetComment(comment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	} else if len(comments) == 0 {
		c.JSON(400, gin.H{"data": nil})
		return
	}
	c.JSON(200, gin.H{"info": "success", "status": 10000, "comments": comments})
}

func AddComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := services.ParseToken(token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user := model.User{
		ID:       claims.ID,
		Username: claims.Username,
	}
	var comment model.Comment
	var product model.Product
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := c.BindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	comment.ProductID = product.ID
	comment.UserID = user.ID
	comment.NickName = user.Username
	err = dao.AddComment(&comment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"info": "success", "status": 10000, "data": comment.ID})
}

func DeleteComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := services.ParseToken(token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	comment := model.Comment{}
	err = c.BindQuery(&comment)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	comment.ID = claims.ID
	err = dao.DeleteComment(comment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"info": "success", "status": 10000})
}

func UpdateComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := services.ParseToken(token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	comment.ID = claims.ID
	err = dao.UpdateComment(&comment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"info": "success", "status": 10000})
}

func Praise(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_, err := services.ParseToken(token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var praise model.Praise
	err = c.BindJSON(&praise)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	err = dao.Praise(praise)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"info": "success", "status": 10000})
}
