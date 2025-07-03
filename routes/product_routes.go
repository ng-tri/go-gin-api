package routes

import (
	"go-gin-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	// productGroup := r.Group("/products", middlewares.AuthMiddleware())
	productGroup := r.Group("/products")
	{
		productGroup.GET("", controllers.GetProducts)
		productGroup.POST("", controllers.CreateProduct)
	}
}
