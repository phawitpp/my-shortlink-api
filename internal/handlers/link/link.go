package linkHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/phawitpp/my-shortlink-api/database"
	"github.com/phawitpp/my-shortlink-api/internal/model"
)

func GetLink(c *fiber.Ctx) error {
	db := database.DB
	var link model.LinkModel

	db.First(&link, "short_link = ?", c.Params("linkUrl"))

	if len(link.ShortLink) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Link not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Link found",
		"data":    link,
	})
}

func CreateLink(c *fiber.Ctx) error {
	db := database.DB
	link := new(model.LinkModel)

	if err := c.BodyParser(link); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	var checkLink model.LinkModel
	db.First(&checkLink, "short_link = ?", link.ShortLink)
	if len(checkLink.ShortLink) > 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Short link already exist",
		})
	}

	err := db.Create(&link).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot create link",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Link created",
		"data":    link,
	})

}

func UpdateLink(c *fiber.Ctx) error {
	db := database.DB
	var link model.LinkModel

	db.First(&link, "short_link = ?", c.Params("linkUrl"))

	if len(link.ShortLink) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Link not found",
		})
	}

	if err := c.BodyParser(&link); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	db.Save(&link)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Link updated",
		"data":    link,
	})
}

func DeleteLink(c *fiber.Ctx) error {
	db := database.DB
	var link model.LinkModel

	db.First(&link, "short_link = ?", c.Params("linkUrl"))

	if len(link.ShortLink) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Link not found",
		})
	}

	db.Delete(&link)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Link deleted",
	})
}
