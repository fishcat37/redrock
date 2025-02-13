package main

import (
	"redrock/api"
	"redrock/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	err := dao.InitSQL()
	if err != nil {
		panic(err)
	}

	err = dao.InitRedis()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	api.Router(r)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
