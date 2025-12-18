package main

import (
	"fmt"
	"go-gin-api/internal/config"
	"go-gin-api/internal/database"
	"go-gin-api/internal/middleware"
	"go-gin-api/internal/route"

	"github.com/gin-gonic/gin"
)

func main() {
	// DATABASE INIT
	config.LoadEnv()
	database.ConnectDB()

	// // AutoMigrate chỉ nên dùng cho demo, project nhỏ hoặc môi trường development
	// database.DB.AutoMigrate(
	// 	&model.User{},
	// 	&model.Product{},
	// 	&model.Order{},
	// 	&model.OrderItem{},
	// )

	// GIN INIT
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middleware.LoggerMiddleware())
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// REGISTER DOMAIN ROUTES
	route.RegisterAuthRoutes(r)
	route.RegisterOrderRoutes(r)
	route.RegisterProductRoutes(r)
	route.RegisterUserRoutes(r)

	// RUN SERVER
	fmt.Println("API Server running at :" + config.Env.AppPort)
	r.Run(":" + config.Env.AppPort)
}
