package rulesets

import (
	"fmt"
	"sort"

	"freecocoa/src/models"
	"freecocoa/src/rulesets/lt75"
	"freecocoa/src/rulesets/lt76"
	"freecocoa/src/rulesets/ltt"
	"freecocoa/src/rulesets/ltx"
)

const LTT = "ltt"
const LTX = "ltx"
const LT75 = "lt75"
const LT76 = "lt76"

func getUnit(unitName, ruleset string) (*models.UnitDetails, error) {
	var unit models.UnitDetails
	var ok bool

	switch ruleset {
	case LTT:
		unit, ok = ltt.UnitStats[unitName]
	case LTX:
		unit, ok = ltx.UnitStats[unitName]
	case LT75:
		unit, ok = lt75.UnitStats[unitName]
	case LT76:
		unit, ok = lt76.UnitStats[unitName]
	default:
		return nil, fmt.Errorf("unknown/unsupported ruleset %s", ruleset)
	}

	if !ok {
		return nil, fmt.Errorf("unit \"%s\" not found", unitName)
	}

	return &unit, nil
}

func getTerrain(terrainName, ruleset string) (*models.TerrainDetails, error) {
	var terrain models.TerrainDetails
	var ok bool

	switch ruleset {
	case LTT:
		terrain, ok = ltt.TerrainStats[terrainName]
	case LTX:
		terrain, ok = ltx.TerrainStats[terrainName]
	case LT75:
		terrain, ok = lt75.TerrainStats[terrainName]
	case LT76:
		terrain, ok = lt76.TerrainStats[terrainName]
	default:
		return nil, fmt.Errorf("unknown/unsupported ruleset %s", ruleset)
	}

	if !ok {
		return nil, fmt.Errorf("terrain \"%s\" not found", terrainName)
	}

	return &terrain, nil
}

func GetAllTerrain(ruleset string) (*[]string, error) {
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
	default:
		return nil, fmt.Errorf("unknown/unsupported ruleset %s", ruleset)
	}

	sort.Strings(knownTerrain)

	return &knownTerrain, nil
}

func GetAllUnits(ruleset string) (*[]string, error) {
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
	default:
		return nil, fmt.Errorf("unknown/unsupported ruleset %s", ruleset)
	}

	sort.Strings(knownUnits)

	return &knownUnits, nil
}
