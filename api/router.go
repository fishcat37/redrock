package api

import (
	"github.com/gin-gonic/gin"
	"redrock/middleware"
	"redrock/service"
)

func Router(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/register", service.Register)
		user.GET("/token", service.Token)
		user.GET("/token/refresh", service.Refresh)
		user.PUT("/password", service.Password)
		user.GET("/info/:user_id", service.GetInfo)
		user.PUT("/info", service.UpdateInfo)
	}
	product := r.Group("/product")
	{
		product.GET("/list", service.List)
		product.GET("/search", service.Search)
		product.PUT("/addCart", service.AddCart)
		product.GET("/cart", service.Cart)
		product.GET("/info/:product_id", service.GetProductInfo)
		product.GET("/type", middleware.AuthMiddleware(), service.GetInfoByType)
	}
	comment := r.Group("/comment")
	{
		comment.GET("/:product_id", service.GetComment)
		comment.POST("/:product_id", service.AddComment)
		comment.DELETE("/:comment_id", service.DeleteComment)
		comment.PUT("/:comment_id", service.UpdateComment)
		comment.PUT("/praise", service.Praise)
	}
	order := r.Group("/operate", middleware.AuthMiddleware())
	{
		order.POST("/order", service.Order)
		order.GET("/:order_id", service.GetOrder)
		order.GET("/list/:user_id", service.GetOrderList)
		order.DELETE("/delete", service.DeleteOrder)
		order.PUT("/update", service.UpdateOrder)
	}

}
