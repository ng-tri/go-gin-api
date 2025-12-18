package handler

import (
	"net/http"
	"strconv"

	"go-gin-api/internal/model"
	"go-gin-api/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{svc: s}
}

func (c *UserHandler) GetUsers(ctx *gin.Context) {
	users, err := c.svc.GetAllUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Không lấy được danh sách user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (c *UserHandler) Register(ctx *gin.Context) {
	var user model.User

	// Lấy JSON từ body
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	// Check email hoặc phone tồn tại (logic trong service)
	if err := c.svc.CheckEmailOrPhoneExists(user.Email, user.Phone); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Gọi service tạo user
	if err := c.svc.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Không thể tạo user",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Tạo user thành công",
		"data":    user,
	})
}

func (c *UserHandler) GetUserByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}

	user, err := c.svc.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *UserHandler) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Query("email")

	user, err := c.svc.GetUserByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy user với email này"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *UserHandler) GetUserByPhone(ctx *gin.Context) {
	phone := ctx.Query("phone")

	user, err := c.svc.GetUserByPhone(phone)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy user với số điện thoại này"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}
