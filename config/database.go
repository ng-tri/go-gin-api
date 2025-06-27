package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// mysql
	// password := os.Getenv("DB_PASSWORD")
	// dsn := fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/go_gin_api?charset=utf8mb4&parseTime=True&loc=Local", password)

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// dsn := os.Getenv("DSN")
	// fmt.Println(dsn)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// postgres
	dsn := "host=localhost user=postgres password=newpassword dbname=go_gin_api port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	DB = db
	fmt.Println("DB connection successful")
}
