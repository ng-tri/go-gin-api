package route

import (
	"go-gin-api/internal/controller"
	"go-gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(r *gin.Engine) {

	authService := service.NewAuthService()
	orderService := service.NewOrderService(authService)
	orderController := controller.NewOrderController(orderService)

	r.POST("/order/create", orderController.CreateOrder)
}
