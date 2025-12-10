package route

import (
	"go-gin-api/internal/controller"
	"go-gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {

	authService := service.NewAuthService()
	authController := controller.NewAuthController(authService)

	auth := r.Group("/auth")

	auth.POST("/login", authController.Login)
	auth.POST("/verify", authController.VerifyToken)
}
