package handlers

import (
	"fiber-postgres-template/database"
	"fiber-postgres-template/middleware"
	"fiber-postgres-template/model"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func CreateNote(c *fiber.Ctx) error {

	user := c.Locals("requestAuth")

	fmt.Println(user)

	fmt.Println(reflect.TypeOf(user))

	// fmt.Println(user.ID)

	details, ok := user.(middleware.DeserializeUser)

	if ok {
		fmt.Println(details.ID)
	} else {
		fmt.Println("raios")
	}

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
