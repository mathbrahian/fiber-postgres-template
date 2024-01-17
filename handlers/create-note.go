package handlers

import uuid "github.com/google/uuid"

import "github.com/gofiber/fiber/v2"

import "fiber-postgres-template/model"

import "fiber-postgres-template/database"

import "fmt"

func CreateNote(c *fiber.Ctx) error {

	request_1 := c.Locals("requestAuth")

	fmt.Println(request_1)
		
	type DeserializeNote struct {
		Note string `json:"note"`
	}

	newNote := new(DeserializeNote)

	if err := c.BodyParser(newNote); err != nil {
		return err
	}

	nota := model.Note{
		// ID: uuid.New(),
		Note: newNote.Note,
		User: uuid.New(),
	}

	db := database.DB
	
	db.Create(&nota)
	
	return c.Status(201).JSON(fiber.Map{"status": "success", "note": nota})
}