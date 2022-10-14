package http

import (
	"freecocoa/src/models"
	"freecocoa/src/utils"

	"github.com/gofiber/fiber/v2"
)

func attack(c *fiber.Ctx) error {
	avd := models.AttackerVDefender{}
	err := c.BodyParser(&avd)
	if err != nil {
		return err
	}

	stats, err := utils.GetStats(avd)
	if err != nil {
		return err
	}

	return c.JSON(stats)
}
