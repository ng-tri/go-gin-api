package route

import (
	"go-gin-api/internal/controller"
	"go-gin-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	// productGroup := r.Group("/products", middlewares.AuthMiddleware())
	// productGroup := r.Group("/products")
	// {
	// 	productGroup.GET("", controller.GetProducts)
	// 	productGroup.POST("", controller.CreateProduct)
	// }

	products := r.Group("/products")
	products.Use(middleware.JWTMiddleware())
	{
		products.GET("", controller.GetProducts)
		products.GET("/:id", controller.GetProduct)
		products.POST("", controller.CreateProduct)
		products.PUT("/:id", controller.UpdateProduct)
		products.DELETE("/:id", controller.DeleteProduct)
	}
}
