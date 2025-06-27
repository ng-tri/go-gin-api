package models

type Product struct {
	// mysql, gorm sẽ tự tạo cột ID, CreatedAt, UpdatedAt, DeletedAt
	// gorm.Model
	// Name  string `json:"name"`
	// Price int    `json:"price"`

	// postgres
	ID    uint   `gorm:"primaryKey" json:"id"`
	Name  string `json:"name" binding:"required,min=2"`
	Price int    `json:"price" binding:"required,gt=0"`
}
