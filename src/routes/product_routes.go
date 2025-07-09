package routes

import (
	"go-gin-api/controllers"
	"go-gin-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	// productGroup := r.Group("/products", middlewares.AuthMiddleware())
	// productGroup := r.Group("/products")
	// {
	// 	productGroup.GET("", controllers.GetProducts)
	// 	productGroup.POST("", controllers.CreateProduct)
	// }

	protected := r.Group("/products")
	protected.Use(middlewares.JWTMiddleware())

	protected.GET("", controllers.GetProducts)
	protected.POST("", controllers.CreateProduct)
	protected.DELETE("/:id", controllers.DeleteProduct)
}
