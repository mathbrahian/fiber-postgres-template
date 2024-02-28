package main

import (
	"fiber-postgres-template/database"
	"fiber-postgres-template/handlers"
	"fiber-postgres-template/model"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

	db := database.DB

	db.AutoMigrate(&model.Sportman{})

	app := fiber.New()

	app.Post("/", handlers.CreateSportman)

	app.Listen(":80")
}
