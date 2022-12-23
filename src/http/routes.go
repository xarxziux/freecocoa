package http

import (
	"fmt"
	"freecocoa/src/core"
	"freecocoa/src/models"

	"github.com/gofiber/fiber/v2"
)

var supportedRulesets []string = []string{"ltt", "ltx", "lt75"}

func attack(c *fiber.Ctx) error {
	ruleset := c.Params("ruleset")

	if !isSupportedRuleset(ruleset) {
		return fmt.Errorf("unknown/unsupported ruleset %s", ruleset)
	}

	avd := models.AttackerVDefender{}
	err := c.BodyParser(&avd)
	if err != nil {
		return err
	}

	baseStats, err := validateInput(avd, ruleset)
	if err != nil {
		return err
	}

	finalStats, err := core.GetStats(*baseStats)
	if err != nil {
		return err
	}

	return c.JSON(finalStats)
}

func isSupportedRuleset(rs string) bool {
	for _, srs := range supportedRulesets {
		if rs == srs {
			return true
		}
	}

	return false
}
