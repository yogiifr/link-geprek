// Package db for database related
package db

import (
	"context"
	"fmt"
	"link-geprek/backend/internal/models"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	RedisClient *redis.Client
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "admin"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_NAME", "shortener"),
		getEnv("DB_PORT", "5432"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Postgres:", err)
	}

	// SAFE MIGRATION
	if !DB.Migrator().HasTable(&models.URL{}) {
		log.Println("Creating table 'urls'...")
		if err := DB.Migrator().CreateTable(&models.URL{}); err != nil {
			log.Fatal("CreateTable failed:", err)
		}
	} else {
		log.Println("Table 'urls' exists, running AutoMigrate...")
		if err := DB.AutoMigrate(&models.URL{}); err != nil {
			if strings.Contains(err.Error(), "does not exist") {
				log.Println("AutoMigrate: constraint issue ignored (safe in dev)")
			} else {
				log.Fatal("AutoMigrate failed:", err)
			}
		}
	}

	// Redis
	redisURL := getEnv("REDIS_URL", "redis://localhost:6379/0")
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatal("Invalid Redis URL:", err)
	}
	RedisClient = redis.NewClient(opt)

	ctx := context.Background()
	if _, err = RedisClient.Ping(ctx).Result(); err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	log.Println("DB & Redis connected!")
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
