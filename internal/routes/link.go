package linkRoutes

import (
	"github.com/gofiber/fiber/v2"
	linkHandler "github.com/phawitpp/my-shortlink-api/internal/handlers/link"
)

func SetupLinkRoutes(router fiber.Router) {
	link := router.Group("/link")

	link.Get("/:linkUrl", linkHandler.GetLink)
	link.Post("/newlink", linkHandler.CreateLink)
	link.Put("/:linkUrl", linkHandler.UpdateLink)
	link.Delete("/:linkUrl", linkHandler.DeleteLink)
}
