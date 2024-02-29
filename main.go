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

	db.AutoMigrate(&model.Sportmen{}, &model.SportmenSport{})

	app := fiber.New()

	app.Post("/", handlers.CreateSportmen)

	app.Listen(":80")
}
