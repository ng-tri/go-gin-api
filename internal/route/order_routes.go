package route

import (
	"go-gin-api/internal/handler"
	"go-gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(r *gin.Engine) {

	authService := service.NewAuthService()
	orderService := service.NewOrderService(authService)
	orderHandler := handler.NewOrderHandler(orderService)

	r.POST("/order/create", orderHandler.CreateOrder)
}
