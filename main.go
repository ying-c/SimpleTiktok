package main

import (
	"app/config"
	"app/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//注册路由
	routes.Setup(r)
	r.Run(fmt.Sprintf(":%d", config.Conf.ServerPort))
}
