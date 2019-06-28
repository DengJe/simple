package main

import (

	"simple/routes"
	"github.com/gin-gonic/gin"



)

func main() {
	app := gin.Default()

	//静态资源
	app.Static("/static", "./static")
	//register routes
	routes.RegisterRoutes(app)
	// 监听并在 0.0.0.0:8080 上启动服务

	app.Run()

}
