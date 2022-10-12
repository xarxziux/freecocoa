package http

import (
	"freecocoa/src/models"
	"freecocoa/src/utils"

	"github.com/gofiber/fiber/v2"
)

func getStats(c *fiber.Ctx) error {
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
}

func attack(c *fiber.Ctx) error {
	return c.JSON("unimplemented")
}
