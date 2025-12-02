package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model

	OrderCode   string  `json:"order_code"`
	TotalAmount float64 `json:"total_amount"`
	Status      string  `json:"status"`

	Items []OrderItem `json:"items"`
}
