package http

import (
	"fmt"
	"sort"

	"freecocoa/src/core"
	"freecocoa/src/models"
	"freecocoa/src/rulesets/lt75"
	"freecocoa/src/rulesets/lt76"
	"freecocoa/src/rulesets/ltt"
	"freecocoa/src/rulesets/ltx"
	"freecocoa/src/warcalc"

	"github.com/gofiber/fiber/v2"
)

var supportedRulesets []string = []string{"ltt", "ltx", "lt75", "lt76"}

const LTT = "ltt"
const LTX = "ltx"
const LT75 = "lt75"
const LT76 = "lt76"

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

	combatResults := warcalc.Warcalc(*finalStats)

	return c.JSON(models.CombinedResults{
		Stats:  *finalStats,
		Combat: combatResults,
	})
}

func calculate(c *fiber.Ctx) error {
	avd := models.FinalStats{}
	err := c.BodyParser(&avd)
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

	combatResults := warcalc.Warcalc(avd)

	return c.JSON(combatResults)
}

func isSupportedRuleset(rs string) bool {
	for _, srs := range supportedRulesets {
		if rs == srs {
			return true
		}
	}

	return false
}

func getTerrain(c *fiber.Ctx) error {
	ruleset := c.Params("ruleset")

	if !isSupportedRuleset(ruleset) {
		return fmt.Errorf("unknown/unsupported ruleset %s", ruleset)
	}

	knownTerrain := make([]string, 0)

	switch ruleset {
	case LTT:
		for terrain := range ltt.TerrainStats {
			knownTerrain = append(knownTerrain, terrain)
		}
	case LTX:
		for terrain := range ltx.TerrainStats {
			knownTerrain = append(knownTerrain, terrain)
		}
	case LT75:
		for terrain := range lt75.TerrainStats {
			knownTerrain = append(knownTerrain, terrain)
		}
	case LT76:
		for terrain := range lt76.TerrainStats {
			knownTerrain = append(knownTerrain, terrain)
		}
	}

	sort.Strings(knownTerrain)

	return c.JSON(knownTerrain)
}

func getUnits(c *fiber.Ctx) error {
	ruleset := c.Params("ruleset")

	if !isSupportedRuleset(ruleset) {
		return fmt.Errorf("unknown/unsupported ruleset %s", ruleset)
	}

	knownUnits := make([]string, 0)

	switch ruleset {
	case LTT:
		for units := range ltt.UnitStats {
			knownUnits = append(knownUnits, units)
		}
	case LTX:
		for units := range ltx.UnitStats {
			knownUnits = append(knownUnits, units)
		}
	case LT75:
		for units := range lt75.UnitStats {
			knownUnits = append(knownUnits, units)
		}
	case LT76:
		for units := range lt76.UnitStats {
			knownUnits = append(knownUnits, units)
		}
	}

	sort.Strings(knownUnits)

	return c.JSON(knownUnits)
}
