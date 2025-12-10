package controller

import (
    "net/http"
    "go-gin-api/internal/model"
    "go-gin-api/internal/service"
    "github.com/gin-gonic/gin"
)

type OrderController struct {
    svc service.OrderService
}

func NewOrderController(s service.OrderService) *OrderController {
    return &OrderController{svc: s}
}

func (oc *OrderController) CreateOrder(c *gin.Context) {
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
