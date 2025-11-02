package server

import (
	"link-geprek/backend/internal/db"
	"link-geprek/backend/internal/handlers"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func Start() {
	db.Init()

	app := fiber.New()

	// Middleware
	app.Use(cors.New())
	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
	}))

	// Routes
	api := app.Group("/api")
	api.Post("/shorten", handlers.ShortenURL)

	app.Get("/:code", handlers.RedirectURL)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on :%s", port)
	log.Fatal(app.Listen(":" + port))
}
