// Package handlers yaa
package handlers

import (
	"fmt"
	"link-geprek/backend/internal/db"
	"link-geprek/backend/internal/models"
	"link-geprek/backend/internal/services"
	"log"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type ShortenRequest struct {
	URL string `JSON:"url" validate:"required,url"`
}

func ShortenURL(c *fiber.Ctx) error {
	var req ShortenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	// validate
	if !govalidator.IsURL(req.URL) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	// normalize
	if !govalidator.IsRequestURL(req.URL) {
		req.URL = "http://" + req.URL
	}

	shortener := services.NewShortener(db.DB)
	code, err := shortener.GenerateShortCode()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate code"})
	}

	url := models.URL{
		ShortCode:   code,
		OriginalURL: req.URL,
	}

	if err := db.DB.Create(&url).Error; err != nil {
		log.Printf("DB INSERT ERROR: %v | URL: %s | Code: %s", err, req.URL, code)
		return c.Status(500).JSON(fiber.Map{"error": "DB error"})
	}

	base := os.Getenv("BASE_URL")
	if base == "" {
		base = "http://localhost:8080"
	}

	return c.JSON(fiber.Map{
		"short_url": fmt.Sprintf("%s/%s", base, code),
		"original":  req.URL,
	})
}
