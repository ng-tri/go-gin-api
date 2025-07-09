package routes

import (
	"go-gin-api/src/controllers"
	"go-gin-api/src/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	// productGroup := r.Group("/products", middlewares.AuthMiddleware())
	// productGroup := r.Group("/products")
	// {
	// 	productGroup.GET("", controllers.GetProducts)
	// 	productGroup.POST("", controllers.CreateProduct)
	// }

	products := r.Group("/products")
	products.Use(middlewares.JWTMiddleware())
	{
		products.GET("", controllers.GetProducts)
		products.GET("/:id", controllers.GetProduct)
		products.POST("", controllers.CreateProduct)
		products.PUT("/:id", controllers.UpdateProduct)
		products.DELETE("/:id", controllers.DeleteProduct)
	}
}
