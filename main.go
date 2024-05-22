package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/phawitpp/my-shortlink-api/database"
	"github.com/phawitpp/my-shortlink-api/internal/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(
		cors.Config{
			AllowOrigins: "http://localhost:3000, https://s.phawit.tech, s.phawit.tech",
    			AllowHeaders:  "Origin, Content-Type, Accept, Authorization",
    			AllowMethods:  "GET, POST, PUT, DELETE, OPTIONS",
    			AllowCredentials: true,
		},
	))
	database.ConnectDB()
	router.SetupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("Hello, World!")
		return err
	})
	app.Listen(":8001")
}
