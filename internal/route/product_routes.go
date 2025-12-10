package route

import (
	"go-gin-api/internal/handler"
	"go-gin-api/internal/middleware"
	"go-gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine) {
	productService := service.NewProductService()
	productHandler := handler.NewProductHandler(productService)

	products := r.Group("/products")
	products.Use(middleware.JWTMiddleware())
	{
		products.GET("", productHandler.GetProducts)
		products.GET("/:id", productHandler.GetProduct)
		products.POST("", productHandler.CreateProduct)
		products.PUT("/:id", productHandler.UpdateProduct)
		products.DELETE("/:id", productHandler.DeleteProduct)
	}
}
