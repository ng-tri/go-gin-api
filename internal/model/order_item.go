package model

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model

	OrderID    uint    `json:"order_id"`    // mã đơn hàng
	ProductID  uint    `json:"product_id"`  // mã sản phẩm
	Quantity   int     `json:"quantity"`    // số lượng trong đơn
	UnitPrice  float64 `json:"unit_price"`  // giá bán tại thời điểm đặt hàng
	TotalPrice float64 `json:"total_price"` // tổng tiền = Quantity * UnitPrice
	Note       string  `json:"note"`        // chú thích (tùy chọn)
	Product    Product `json:"product"`
}
