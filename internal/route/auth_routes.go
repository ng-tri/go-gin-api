package route

import (
	"go-gin-api/internal/handler"
	"go-gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {

	authService := service.NewAuthService()
	authHandler := handler.NewAuthHandler(authService)

	auth := r.Group("/auth")

	auth.POST("/login", authHandler.Login)
	auth.POST("/verify", authHandler.VerifyToken)
}
