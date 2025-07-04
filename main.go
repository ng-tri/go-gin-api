package main

import (
	"go-gin-api/config"
	"go-gin-api/middlewares"
	"go-gin-api/models"
	"go-gin-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Product{})
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middlewares.LoggerMiddleware())
	r.SetTrustedProxies([]string{"127.0.0.1"})

	routes.RegisterAuthRoutes(r)
	routes.RegisterProductRoutes(r)

	r.Run(":8080")
}
