package http

import (
	"fmt"
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
	app.Post("/api/v1/:ruleset/buildCost", getBuildCost)

	logo := "______\n" +
		"|  ___|\n" +
		"| |_ _ __ ___  ___  ___ ___   ___ ___   __ _\n" +
		"|  _| '__/ _ \\/ _ \\/ __/ _ \\ / __/ _ \\ / _` |\n" +
		"| | | | |  __/  __/ (_| (_) | (_| (_) | (_| |\n" +
		"\\_| |_|  \\___|\\___|\\___\\___/ \\___\\___/ \\__,_|\n" +
		"\n" +
		"     (  )   (   )  )\n" +
		"      ) (   )  (  (\n" +
		"      ( )  (    ) )\n" +
		"      _____________\n" +
		"     <_____________> ___\n" +
		"     |             |/ _ \\\n" +
		"     |               | | |\n" +
		"     |               |_| |\n" +
		"  ___|             |\\___/\n" +
		" /    \\___________/    \\\n" +
		" \\_____________________/\n"

	fmt.Println(logo)

	log.Fatal(app.Listen(":7123"))
}
