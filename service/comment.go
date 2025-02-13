package service

import (
	"net/http"
	"redrock/config"
	"redrock/dao"
	"redrock/model"
	"redrock/utils"

	"github.com/gin-gonic/gin"
)

func GetComment(c *gin.Context) {
	var comment model.Comment
	if err := c.BindUri(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   err.Error()})
		return
	}
	comments, err := dao.GetComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": config.DataBaseErrCode,
			"info":   err.Error()})
		return
	} else if len(comments) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.NotCommentsCode,
			"data":   nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": config.SuccessCode, "data": comments})
}

func AddComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.TokenErrCode,
			"info":   err.Error()})
		return
	}
	user := model.User{
		ID:       claims.ID,
		Username: claims.Username,
	}
	var comment model.Comment
	//var product model.Product
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   err.Error()})
		return
	}
	//if err := c.BindJSON(&product); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"info": err.Error()})
	//	return
	//}
	//comment.ProductID = product.ID
	comment.UserID = user.ID
	comment.NickName = user.Username
	err = dao.AddComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": config.DataBaseErrCode,
			"info":   err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"info": "success", "status": config.SuccessCode, "data": gin.H{"comment_ID": comment.ID}})
}

func DeleteComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.TokenErrCode,
			"info":   err.Error()})
		return
	}
	comment := model.Comment{}
	err = c.BindUri(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   err.Error()})
		return
	}
	comment.UserID = claims.ID
	err = dao.DeleteComment(comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.DataBaseErrCode,
			"info":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": config.SuccessCode})
}

func UpdateComment(c *gin.Context) {
	token := c.GetHeader("Authorization")
	claims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.TokenErrCode,
			"info":   err.Error()})
		return
	}
	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   err.Error()})
		return
	}
	comment.UserID = claims.ID
	err = dao.UpdateComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": config.DataBaseErrCode,
			"info":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": config.SuccessCode})
}

func Praise(c *gin.Context) {
	token := c.GetHeader("Authorization")
	_, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.TokenErrCode,
			"info":   err.Error()})
		return
	}
	var praise model.Praise
	err = c.Bind(&praise)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": config.RequestErrCode,
			"info":   err.Error()})
		return
	}
	err = dao.Praise(praise)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": config.DataBaseErrCode,
			"info":   err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"info": "success", "status": config.SuccessCode})
}
