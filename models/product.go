package models

type Product struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name" binding:"required,min=2"`
	Price int    `json:"price" binding:"required,gt=0"`
}
