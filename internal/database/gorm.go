package database

import (
	"fmt"
	"go-gin-api/internal/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := config.Env.DatabaseURL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	DB = db
	fmt.Println("DB connection successful")
}
