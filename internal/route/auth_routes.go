package route

import (
	"go-gin-api/internal/controller"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	r.POST("/login", controller.Login)
}
