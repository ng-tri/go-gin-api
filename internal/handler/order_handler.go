package handler

import (
	"go-gin-api/internal/model"
	"go-gin-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	svc service.OrderService
}

func NewOrderHandler(s service.OrderService) *OrderHandler {
	return &OrderHandler{svc: s}
}

func (oc *OrderHandler) CreateOrder(c *gin.Context) {
	var order model.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order data"})
		return
	}

	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		return
	}

	// Gọi service xử lý logic
	if err := oc.svc.CreateOrder(order, token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order created",
		"order":   order,
	})
}
