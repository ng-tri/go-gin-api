package handler

import (
	"net/http"
	"strconv"

	"go-gin-api/internal/model"
	"go-gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	svc *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{svc: s}
}

func (c *ProductHandler) GetProducts(ctx *gin.Context) {
	products, err := c.svc.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (c *ProductHandler) GetProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	product, err := c.svc.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (c *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req model.Product
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	p, err := c.svc.Create(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, p)
}

func (c *ProductHandler) UpdateProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var req model.Product
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	p, err := c.svc.Update(uint(id), &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, p)
}

func (c *ProductHandler) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	err := c.svc.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
