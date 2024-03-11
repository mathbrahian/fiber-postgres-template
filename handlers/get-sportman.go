package handlers

import (
	"fiber-postgres-template/database"
	"fiber-postgres-template/middleware"
	"fiber-postgres-template/model"

	"github.com/gofiber/fiber/v2"
)

func GetSport(c *fiber.Ctx) error {

	user := c.Locals("requestAuth")

	details, _ := user.(middleware.DeserializeUser)

	sport := model.Sportmen{
		UserId: details.ID,
	}

	db := database.DB

	db.First(&sport)

	return c.Status(201).JSON(sport)
}
