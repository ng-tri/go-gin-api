package migrate

import (
	"fmt"
	"log"
	"os"

	"go-gin-api/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DSN") // ví dụ: "user:pass@tcp(localhost:3306)/go_gin_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}

	fmt.Println("Running migration...")
	err = db.AutoMigrate(
		&model.User{},
		// thêm nhiều model nếu cần
	)
	if err != nil {
		log.Fatal("migration failed:", err)
	}

	fmt.Println("Migration completed successfully!")
}
