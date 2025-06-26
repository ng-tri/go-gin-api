package routes

import (
	"go_gin_api/controllers"
	"go_gin_api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	productGroup := r.Group("/products", middlewares.AuthMiddleware())
	{
		productGroup.GET("", controllers.GetProducts)
		productGroup.POST("", controllers.CreateProduct)
	}
}
