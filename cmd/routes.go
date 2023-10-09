package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Maxibanex/prueba_paralelo.git/handlers"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", handlers.ListFacts)

	app.Get("/fact", handlers.NewFactView) // Add new route for new view
	app.Post("/fact", handlers.CreateFact)

}
