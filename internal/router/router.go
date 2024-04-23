package router

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	linkRoutes "github.com/phawitpp/my-shortlink-api/internal/routes"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", func(c *fiber.Ctx) error {
		log.Println("Time:", time.Now())
		return c.Next()
	})

	linkRoutes.SetupLinkRoutes(api)
}
