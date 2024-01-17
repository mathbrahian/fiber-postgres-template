package main

import "github.com/gofiber/fiber/v2"
import "fiber-postgres-template/database"
import "fiber-postgres-template/handlers"
import "fiber-postgres-template/model"
import "fiber-postgres-template/middleware"

// Custom middleware function
func Piripiropi(c *fiber.Ctx) error {
    // Perform tasks before the route handling function
    println("Middleware: Request received")
    
    // Continue to the next middleware or route handling function
    return c.Next()
}


func main() {
	database.ConnectDB()

	db := database.DB
	db.AutoMigrate(&model.Note{})

	app := fiber.New()

	app.Post("/hi", handlers.CreateNote)
	
	app.Use(Piripiropi)

	app.Use(middleware.New(middleware.Config{}))

	app.Post("/", handlers.CreateNote)

	app.Listen(":80")
}
