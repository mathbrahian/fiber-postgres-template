package handlers

import (
	"fiber-postgres-template/database"
	"fiber-postgres-template/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateSportmen(c *fiber.Ctx) error {

	Sportmen := new(model.Sportmen)

	if err := c.BodyParser(Sportmen); err != nil {
		return err
	}

	var validate *validator.Validate

	validate = validator.New(validator.WithRequiredStructEnabled())

	err1 := validate.Struct(Sportmen)

	err2 := validate.Struct(Sportmen.Sport)

	if err1 != nil && err2 != nil {

		return c.Status(400).JSON(fiber.Map{"status": "error"})

	}

	db := database.DB

	db.Create(&Sportmen)

	return c.Status(201).JSON(fiber.Map{"status": "success", "sportman": Sportmen})
}
