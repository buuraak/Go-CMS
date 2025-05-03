package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env variables")
	}
}

func GetPort() string {
	return GetEnv("PORT", DefaultPort)
}

func GetJWTSecret() string {
	secret := GetEnv("JWT_SECRET", "")
	if secret == "" {
		log.Fatal("JWT_SECRET is required but not set")
	}
	return secret
}

func GetDbDsn() string {
	dsn := GetEnv("DB_DSN", "")
	if dsn == "" {
		log.Fatal("DB_DSN is required but not set")
	}
	return dsn
}

func GetEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
