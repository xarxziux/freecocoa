package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type unit struct {
	Unit string `json:"unit"`
}

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		if c.Is("json") {
			return c.Next()
		}
		return c.SendString("Only JSON allowed!")
	})

	baseUnit := unit{}

	app.Post("/api/v0.0.1/getdefence", func(c *fiber.Ctx) error {
		err := c.BodyParser(&baseUnit)

		if err != nil {
			return c.SendString(":(")
		}

		return c.JSON(baseUnit)
		//return c
	})

	log.Fatal(app.Listen(":7123"))
}
