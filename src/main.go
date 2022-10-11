package main

import (
	"log"

	"freecocoa/src/models"
	"freecocoa/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

/*type unit struct {
	Unit string `json:"unit"`
        }*/

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		if c.Is("json") {
			return c.Next()
		}
		return c.SendString("Only JSON allowed!")
	})

	app.Post("/api/v0.0.1/getstats", func(c *fiber.Ctx) error {
		baseUnit := models.DefenderUnit{}
		err := c.BodyParser(&baseUnit)

		if err != nil {
			return err
		}

		fullUnit, err := utils.ConvertDefender(baseUnit)
		if err != nil {
			return err
		}

		return c.JSON(fullUnit)
	})

	log.Fatal(app.Listen(":7123"))
}
