package main

import (
	"backend/backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.NewHandler()
	log.Fatal(app.Listen(":3000"))
}
