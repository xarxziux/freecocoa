package http

import (
	"freecocoa/src/core"
	"freecocoa/src/models"
	"freecocoa/src/rulesets"
	"freecocoa/src/warcalc"

	"github.com/gofiber/fiber/v2"
)

func attack(c *fiber.Ctx) error {
	ruleset := c.Params("ruleset")
	ai := &models.AttackInput{}
	err := c.BodyParser(ai)
	if err != nil {
		return err
	}

	validatedStats, err := rulesets.PopulateInput(ai, ruleset)
	if err != nil {
		return err
	}

	attackResults := core.GetStats(validatedStats)
	combatResults := warcalc.Warcalc(attackResults)

	return c.JSON(models.CombinedResults{
		Stats:  attackResults,
		Combat: combatResults,
	})
}

func calculate(c *fiber.Ctx) error {
	ar := models.AttackResults{}
	err := c.BodyParser(&ar)
	if err != nil {
		return err
	}

	// TODO: Add validation
	/*
		baseStats, err := validateInput(avd, ruleset)
		if err != nil {
			return err
		}
	*/

	combatResults := warcalc.Warcalc(&ar)

	return c.JSON(combatResults)
}

func getBuildCost(c *fiber.Ctx) error {
	ruleset := c.Params("ruleset")
	bci := &models.BuildCostInput{}
	err := c.BodyParser(bci)
	if err != nil {
		return err
	}

	err = rulesets.PopulateBuildCost(bci, ruleset)
	if err != nil {
		return err
	}

	costs := core.GetBuildCosts(bci)

	return c.JSON(costs)
}

func getTerrain(c *fiber.Ctx) error {
	ruleset := c.Params("ruleset")
	knownTerrain, err := rulesets.GetAllTerrain(ruleset)

	if err != nil {
		return err
	}

	return c.JSON(knownTerrain)
}

func getUnits(c *fiber.Ctx) error {
	ruleset := c.Params("ruleset")
	knownUnits, err := rulesets.GetAllUnits(ruleset)

	if err != nil {
		return err
	}

	return c.JSON(knownUnits)
}
