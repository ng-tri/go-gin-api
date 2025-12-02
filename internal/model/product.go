package model

import "gorm.io/gorm"

type Product struct {
	// mysql, gorm sẽ tự tạo cột ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`

	// postgres
	// ID          uint   `gorm:"primaryKey" json:"id"`
	// Name        string `json:"name" binding:"required,min=2"`
	// Price       int    `json:"price" binding:"required,gt=0"`
	// Description string `json:"description"`
	// Stock       int    `json:"stock" binding:"required,gt=0"`
}
