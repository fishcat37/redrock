package api

import (
	"github.com/gin-gonic/gin"
	"redrock/service"
)

func Router(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/register", service.Register)
		user.POST("/token", service.Token)
		user.GET("/token/refresh", service.Refresh)
		user.PUT("/password", service.Password)
		user.GET("/info/:user_id", service.GetInfo)
		user.PUT("/info", service.UpdateInfo)
	}
	product := r.Group("/product")
	{
		product.GET("/list", service.List)
		product.GET("/search", service.Search)
		product.PUT("/AddCart", service.AddCart)
		product.GET("/cart", service.Cart)
		product.GET("/info/:product_id", service.GetProductInfo)
		product.GET("/:type", service.GetInfoByType)
	}
	comment := r.Group("/comment")
	{
		comment.GET("/:product_id", service.GetComment)
		comment.POST("/:product_id", service.AddComment)
		comment.DELETE("/:comment_id", service.DeleteComment)
		comment.PUT("/:comment_id", service.UpdateComment)
		comment.PUT("/praise", service.Praise)
	}
	r.POST("operate/order", service.Order)
}
