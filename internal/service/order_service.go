package service

import (
	"encoding/json"
	"errors"
	"go-gin-api/internal/model"
	"io"
	"net/http"
)

type Order struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type OrderService interface {
	CreateOrder(order model.Order, token string) error
}

type orderService struct{}

func NewOrderService(auth AuthService) OrderService {
	return &orderService{}
}

func (s *orderService) CreateOrder(order model.Order, token string) error {

	claims, err := verifyToken(token)
	if err != nil {
		return errors.New("invalid token")
	}

	// Xử lý logic nghiệp vụ
	userID := claims["userID"]
	if userID == nil {
		return errors.New("missing user ID in token")
	}

	// TODO: Insert DB tại đây

	return nil
}

func verifyToken(token string) (map[string]interface{}, error) {
	client := &http.Client{}
	// req, err := http.NewRequest("POST", "http://localhost:8081/auth/verify", nil)
	req, err := http.NewRequest("POST", "http://auth:8081/auth/verify", nil) // Nếu chạy trong docker-compose, dùng `http://auth:8081`
	req.Header.Set("Authorization", token)                                   // token đã có prefix "Bearer "
	if err != nil {
		return nil, err
	}

	// req.Header.Add("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var authResponse map[string]interface{}
	json.Unmarshal(body, &authResponse)
	return authResponse, nil
}
