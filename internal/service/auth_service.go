package service

import (
	"errors"
	"os"
	"time"

	"go-gin-api/internal/database"
	"go-gin-api/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthService interface {
	Login(req LoginRequest) (string, error)
	VerifyToken(tokenStr string) (any, error)
}

type authService struct{}

func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) Login(req LoginRequest) (string, error) {
	// Lấy user theo username
	var user model.User

	err := database.DB.
		Where("username = ?", req.Username).
		First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("username không tồn tại")
		}
		return "", err
	}

	// Kiểm tra password (ở đây là plain password – nếu hash thì mình hỗ trợ thêm)
	if user.Password != req.Password {
		return "", errors.New("sai mật khẩu")
	}

	// Tạo claims JWT
	claims := jwt.MapClaims{
		"username": user.Email,
		"user_id":  user.ID,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	// Tạo token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("không tạo được token")
	}

	return tokenStr, nil
}

func (s *authService) VerifyToken(tokenStr string) (any, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token.Claims, nil
}
