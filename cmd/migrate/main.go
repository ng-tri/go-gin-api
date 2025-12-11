package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

var migrationDir = "./internal/migration"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: migrate [migrate|rollback|fresh|status]")
	}

	command := os.Args[1]

	godotenv.Load()
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("Missing DSN environment variable")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB connect error:", err)
	}

	goose.SetDialect("mysql")

	switch command {
	case "migrate":
		migrateWithBatch(db)

	case "rollback":
		rollbackBatch(db)

	case "fresh":
		freshDatabase(db)

	case "status":
		goose.Status(db, migrationDir)

	default:
		log.Fatal("Unknown command:", command)
	}
}

func migrateWithBatch(db *sql.DB) {
	fmt.Println("Running migrations with batch...")

	// 1. Lấy migration chưa chạy
	pending, err := collectPending(db, migrationDir)
	if err != nil {
		log.Fatal("Error collecting pending migrations:", err)
	}

	if len(pending) == 0 {
		fmt.Println("No pending migrations")
		return
	}

	// 2. Lấy batch hiện tại
	batch := getLatestBatch(db) + 1

	// 3. Chạy từng migration
	for _, m := range pending {
		fmt.Println("Applying:", m.Version)

		err := goose.UpTo(db, migrationDir, m.Version)
		if err != nil {
			log.Fatal(err)
		}

		// 4. Ghi vào migration_batches
		_, err = db.Exec(
			`INSERT INTO migration_batches (version_id, batch, migration) VALUES (?, ?, ?)`,
			m.Version, batch, m.Source,
		)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Migration complete (batch:", batch, ")")
}

func rollbackBatch(db *sql.DB) {
	fmt.Println("Rolling back latest batch...")

	batch := getLatestBatch(db)
	if batch == 0 {
		fmt.Println("Nothing to rollback")
		return
	}

	// lấy danh sách migration của batch cuối
	rows, err := db.Query(
		`SELECT version_id FROM migration_batches WHERE batch = ? ORDER BY id DESC`,
		batch,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var versions []int64
	for rows.Next() {
		var v int64
		rows.Scan(&v)
		versions = append(versions, v)
	}

	// rollback từng version
	for _, v := range versions {
		fmt.Println("Rollback:", v)

		// DownTo(v-1) = rollback đúng migration v
		err := goose.DownTo(db, migrationDir, v-1)
		if err != nil {
			log.Fatal(err)
		}
	}

	// xóa batch trong DB
	_, _ = db.Exec(`DELETE FROM migration_batches WHERE batch = ?`, batch)

	fmt.Println("Rollback batch", batch, "done")
}

func freshDatabase(db *sql.DB) {
	fmt.Println("Resetting database...")

	// rollback toàn bộ migration
	err := goose.Reset(db, migrationDir)
	if err != nil {
		log.Fatal(err)
	}

	// xóa bảng batch
	_, _ = db.Exec(`TRUNCATE TABLE migration_batches`)

	// migrate lại từ đầu
	migrateWithBatch(db)

	fmt.Println("Database is fresh!")
}

func getLatestBatch(db *sql.DB) int {
	var batch sql.NullInt64
	db.QueryRow(`SELECT MAX(batch) FROM migration_batches`).Scan(&batch)

	if batch.Valid {
		return int(batch.Int64)
	}
	return 0
}

func collectPending(db *sql.DB, dir string) ([]*goose.Migration, error) {
	// 1. phiên bản hiện tại của DB
	current, err := goose.GetDBVersion(db)
	if err != nil {
		return nil, err
	}

	// 2. load toàn bộ migration
	all, err := goose.CollectMigrations(dir, 0, goose.MaxVersion)
	if err != nil {
		return nil, err
	}

	// 3. lọc migration version > current
	pending := []*goose.Migration{}
	for _, m := range all {
		if m.Version > current {
			pending = append(pending, m)
		}
	}

	return pending, nil
}
