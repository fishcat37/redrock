package api

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/register", Register)
		user.POST("/token", Token)
		user.GET("/token/refresh", Refresh)
		user.PUT("/password", Password)
		user.GET("info/:user_id", GetInfo)
		user.PUT("/info", UpdateInfo)
	}
	product := r.Group("/product")
	{
		product.GET("/list", List)
		product.GET("/search", Search)
		product.PUT("/AddCart", AddCart)
		product.GET("/cart", Cart)
		product.GET("/info/:product_id", GetProductInfo)
	}
	comment := r.Group("/comment")
	{
		comment.GET("/:product_id", GetComment)
		comment.POST("/:product_id", AddComment)
		comment.DELETE("/:comment_id", DeleteComment)
		comment.PUT(":comment_id", UpdateComment)
		comment.PUT("/praise", Praise)
	}
	r.POST("operate/order", Order)
}
