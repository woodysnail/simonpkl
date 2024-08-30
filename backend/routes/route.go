package routes

import "github.com/gofiber/fiber/v2"

func NewHandler() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

}
