package handlers

import (
	"fiber-postgres-template/database"
	"fiber-postgres-template/middleware"
	"fiber-postgres-template/model"

	"github.com/gofiber/fiber/v2"
)

func CreateNote(c *fiber.Ctx) error {

	user := c.Locals("requestAuth")

	details, _ := user.(middleware.DeserializeUser)

	type DeserializeNote struct {
		Note string `json:"note"`
	}

	newNote := new(DeserializeNote)

	if err := c.BodyParser(newNote); err != nil {
		return err
	}

	nota := model.Note{
		Note: newNote.Note,
		User: details.ID,
	}

	db := database.DB

	db.Create(&nota)

	return c.Status(201).JSON(fiber.Map{"status": "success", "note": nota})
}
