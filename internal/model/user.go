package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex;size:255"`
	Phone    string `gorm:"uniqueIndex;size:20"`
	Password string `json:"-"`
}
