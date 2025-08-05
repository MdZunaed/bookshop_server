package main

import (
	"fmt"

	"github.com/MdZunaed/bookshop/config"
	"github.com/MdZunaed/bookshop/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.Use(gin.Recovery())
	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
			"message": "Server is running fine",
		})
	})
	routes.RegisterRoute(app)

	fmt.Println("Server started")

	port := config.GetEnvProperty("port")
	app.Run(fmt.Sprintf(":%s", port))
}
