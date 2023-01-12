package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func StartServer() {
	app := fiber.New()
	app.Use(cors.New())

	/*app.Use(func(c *fiber.Ctx) error {
		if c.Is("json") {
			return c.Next()
		}
		return c.SendString("Only JSON allowed!")
	})*/

	app.Post("/api/v1/:ruleset/attack", attack)
	app.Get("/api/v1/:ruleset/getTerrain", getTerrain)
	app.Get("/api/v1/:ruleset/getUnits", getUnits)
	app.Post("/api/v1/calculate", calculate)

	log.Fatal(app.Listen(":7123"))
}
