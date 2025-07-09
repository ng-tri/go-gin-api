package controllers

import (
	"fmt"
	"go-gin-api/src/config"
	"go-gin-api/src/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&newProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create"})
		return
	}

	c.JSON(http.StatusOK, newProduct)
}

func DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	fmt.Println(string(idStr))
	fmt.Scanf(string(idStr))
	os.Exit(0)
	id, err := strconv.Atoi(idStr)
	log.Printf("%d\n", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product models.Product
	result := config.DB.First(&product, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
