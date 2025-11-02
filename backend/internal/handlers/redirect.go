package handlers

import (
	"context"
	"link-geprek/backend/internal/db"
	"link-geprek/backend/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RedirectURL(c *fiber.Ctx) error {
	code := c.Params("code")
	ctx := context.Background()

	// 1. Try Redis cache
	cached, err := db.RedisClient.Get(ctx, "url:"+code).Result()
	if err == nil && cached != "" {
		go incrementClickAsync(code) // fire & forget
		return c.Redirect(cached, 301)
	}

	// 2. Fallback to DB
	var url models.URL
	if err := db.DB.Where("short_code = ?", code).First(&url).Error; err != nil {
		return c.Status(404).SendString("Short URL not found")
	}

	// 3. Cache to Redis (1 hour)
	db.RedisClient.Set(ctx, "url:"+code, url.OriginalURL, 1*time.Hour)

	// 4. Increment click
	go incrementClickAsync(code)

	return c.Redirect(url.OriginalURL, 301)
}

func incrementClickAsync(code string) {
	ctx := context.Background()
	// Try Redis INCR
	key := "clicks:" + code
	if _, err := db.RedisClient.Incr(ctx, key).Result(); err == nil {
		return
	}

	// Fallback: DB
	db.DB.Model(&models.URL{}).Where("short_code = ?", code).UpdateColumn("clicks", gorm.Expr("clicks + 1"))
}
