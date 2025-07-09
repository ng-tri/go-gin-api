package routes

import (
	"go-gin-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	r.POST("/login", controllers.Login)
}
