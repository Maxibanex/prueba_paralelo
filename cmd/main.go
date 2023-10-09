package main

import (
	"github.com/Maxibanex/prueba_paralelo.git/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Maxibanex rodriguez!")
	})

	app.Listen(":3000")
}
