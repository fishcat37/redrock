package main

import (
	"redrock/api"
	"redrock/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	// 3. 初始化数据库
	err := dao.InitDB()
	if err != nil {
		panic(err)
	}
	// 4. 注册路由
	r := gin.Default()
	// 5. 启动服务
	api.Router(r)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
