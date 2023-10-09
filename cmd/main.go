package main

import (
	"github.com/Maxibanex/prueba_paralelo.git/database"
	"github.com/Maxibanex/prueba_paralelo.git/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,         // new config
		ViewsLayout: "layouts/main", // add this to config
	})

	setupRoutes(app)

	app.Static("/", "./public")

	// Set up 404 page
	app.Use(handlers.NotFound)

	app.Listen(":3000")
}
