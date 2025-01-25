package api

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.POST("/user/register", Register)
	r.POST("/user/token", Token)
	r.GET("/user/token/refresh", Refresh)
	r.POST("/user/password", Password)
	r.GET("/user/info/:user_id", GetInfo)
	r.POST("/user/info", UpdateInfo)
}
