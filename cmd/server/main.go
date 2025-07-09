package main

import (
	"go-gin-api/internal/database"
	"go-gin-api/internal/middleware"
	"go-gin-api/internal/model"
	"go-gin-api/internal/route"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	database.DB.AutoMigrate(&model.Product{})
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middleware.LoggerMiddleware())
	r.SetTrustedProxies([]string{"127.0.0.1"})

	route.RegisterAuthRoutes(r)
	route.RegisterProductRoutes(r)

	r.Run(":8080")
}
