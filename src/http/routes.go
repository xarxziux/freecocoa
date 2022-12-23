package http

import (
	"freecocoa/src/core"
	"freecocoa/src/models"

	"github.com/gofiber/fiber/v2"
)

func attack(c *fiber.Ctx) error {
	avd := models.AttackerVDefender{}
	err := c.BodyParser(&avd)
	if err != nil {
		return err
	}

	baseStats, err := validateInput(avd)
	if err != nil {
		return err
	}

	finalStats, err := core.GetStats(*baseStats)
	if err != nil {
		return err
	}

	return c.JSON(finalStats)
}
