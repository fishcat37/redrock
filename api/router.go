package api

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.POST("/user/register", Register)
	r.POST("/user/token", Token)
	r.GET("/user/token/refresh", Refresh)
	r.PUT("/user/password", Password)
	r.GET("/user/info/:user_id", GetInfo)
	r.PUT("/user/info", UpdateInfo)
	r.GET("/product/list", List)
	r.GET("/book/search", Search)
	r.PUT("/product/AddCart", AddCart)
	r.GET("product/cart", Cart)
}
