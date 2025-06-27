package main

import (
	"fmt"
	"go_gin_api/config"
	"go_gin_api/models"
	"go_gin_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	fmt.Println("DB connection successful")
	config.DB.AutoMigrate(&models.Product{})

	r := gin.Default()

	routes.RegisterProductRoutes(r)

	r.Run(":8080")
}
