package route

import (
	"go-gin-api/internal/controller"
	"go-gin-api/internal/middleware"
	"go-gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	productService := service.NewProductService()
	productController := controller.NewProductController(productService)

	products := r.Group("/products")
	products.Use(middleware.JWTMiddleware())
	{
		products.GET("", productController.GetProducts)
		products.GET("/:id", productController.GetProduct)
		products.POST("", productController.CreateProduct)
		products.PUT("/:id", productController.UpdateProduct)
		products.DELETE("/:id", productController.DeleteProduct)
	}
}
