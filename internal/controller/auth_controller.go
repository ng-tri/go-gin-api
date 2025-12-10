package controller

import (
	"net/http"

	"go-gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	svc service.AuthService
}

func NewAuthController(s service.AuthService) *AuthController {
	return &AuthController{svc: s}
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req service.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	token, err := c.svc.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *AuthController) VerifyToken(ctx *gin.Context) {
	token := ctx.Query("token")

	result, err := c.svc.VerifyToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
