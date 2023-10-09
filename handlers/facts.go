package handlers

import (
	"github.com/Maxibanex/prueba_paralelo.git/database"
	"github.com/Maxibanex/prueba_paralelo.git/models"
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)

	return c.Render("index", fiber.Map{
		"Title":    "Diversion",
		"Subtitle": "Facts for funtimes with friends!",
		"Facts":    facts, // send the facts to the view
	})
}

// Create new Fact View handler
func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Fact",
		"Subtitle": "Add a cool fact!",
	})
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	// Parse request body
	if err := c.BodyParser(fact); err != nil {
		return NewFactView(c)
	}

	// Create fact in database
	result := database.DB.Db.Create(&fact)
	if result.Error != nil {
		return NewFactView(c)
	}

	return ConfirmationView(c) // 2. Return confirmation view
}

// 1. New Confirmation view
func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"Title":    "Fact added successfully",
		"Subtitle": "Add more wonderful facts to the list!",
	})
}

func ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&fact)
	if result.Error != nil {
		return NotFound(c)
	}

	return c.Status(fiber.StatusOK).Render("show", fiber.Map{
		"Title": "Single Fact",
		"Fact":  fact,
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
}

func EditFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).First(&fact)
	if result.Error != nil {
		return NotFound(c)
	}

	return c.Render("edit", fiber.Map{
		"Title":    "Edit Fact",
		"Subtitle": "Edit your interesting fact",
		"Fact":     fact,
	})
}

func UpdateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	id := c.Params("id")

	// Parsing the request body
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}

	// Write updated values to the database
	result := database.DB.Db.Model(&fact).Where("id = ?", id).Updates(fact)
	if result.Error != nil {
		return EditFact(c)
	}

	return ShowFact(c)
}

func DeleteFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := database.DB.Db.Where("id = ?", id).Delete(&fact)
	if result.Error != nil {
		return NotFound(c)
	}

	return ListFacts(c)
}
