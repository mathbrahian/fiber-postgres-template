package main

import (
	"fiber-postgres-template/database"
	"fiber-postgres-template/handlers"
	"fiber-postgres-template/middleware"
	"fiber-postgres-template/model"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

	db := database.DB

	db.AutoMigrate(&model.Note{})

	app := fiber.New()

	app.Use(middleware.New(middleware.Config{}))

	app.Post("/", handlers.CreateNote)

	app.Listen(":80")
}
