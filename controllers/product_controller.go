package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "go_gin_api/models"
)

var products = []models.Product{
    {ID: 1, Name: "iPhone", Price: 25000000},
    {ID: 2, Name: "MacBook", Price: 45000000},
}

func GetProducts(c *gin.Context) {
    c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
    var newProduct models.Product
    if err := c.ShouldBindJSON(&newProduct); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newProduct.ID = len(products) + 1
    products = append(products, newProduct)
    c.JSON(http.StatusOK, newProduct)
}
