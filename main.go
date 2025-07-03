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

	r := gin.Default()
	r.Use(middlewares.Logger())

	routes.RegisterProductRoutes(r)

	r.Run(":8080")
}
