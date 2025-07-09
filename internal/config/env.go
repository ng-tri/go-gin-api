package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type EnvConfig struct {
    AppPort      string
    DatabaseURL  string
    JwtSecret    string
    RedisAddress string
}

var Env EnvConfig

func LoadEnv() {
    // Load từ file .env nếu có
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables")
    }

    Env = EnvConfig{
        AppPort:      getEnv("APP_PORT", "8080"),
        DatabaseURL:  getEnv("DATABASE_URL", ""),
        JwtSecret:    getEnv("JWT_SECRET", "default-secret"),
        RedisAddress: getEnv("REDIS_ADDRESS", "localhost:6379"),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
