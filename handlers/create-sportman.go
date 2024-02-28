package handlers

import (
	"fiber-postgres-template/database"
	"fiber-postgres-template/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateSportman(c *fiber.Ctx) error {

	Sportman := new(model.Sportman)

	if err := c.BodyParser(Sportman); err != nil {
		return err
	}

	var validate *validator.Validate

	validate = validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(Sportman)

	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		// if _, ok := err.(*validator.InvalidValidationError); ok {
		// 	fmt.Println("aja")
		// 	fmt.Println(err)
		// }

		// for _, err := range err.(validator.ValidationErrors) {

		// 	fmt.Println(err.Field())
		// }
		return c.Status(400).JSON(fiber.Map{"status": "error"})

	}

	db := database.DB

	db.Create(&Sportman)

	return c.Status(201).JSON(fiber.Map{"status": "success", "sportman": Sportman})
}
