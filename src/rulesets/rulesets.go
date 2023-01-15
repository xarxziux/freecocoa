package rulesets

import (
	"fmt"

	"freecocoa/src/models"
	"freecocoa/src/rulesets/lt75"
	"freecocoa/src/rulesets/lt76"
	"freecocoa/src/rulesets/ltt"
	"freecocoa/src/rulesets/ltx"
)

func GetUnit(unitName, ruleset string) (*models.UnitDetails, error) {
	var unit models.UnitDetails
	var ok bool

	switch ruleset {
	case "ltt":
		unit, ok = ltt.UnitStats[unitName]
	case "ltx":
		unit, ok = ltx.UnitStats[unitName]
	case "lt75":
		unit, ok = lt75.UnitStats[unitName]
	case "lt76":
		unit, ok = lt76.UnitStats[unitName]
	default:
		return nil, fmt.Errorf("invalid ruleset")
	}

	if !ok {
		return nil, fmt.Errorf("unit \"%s\" not found", unitName)
	}

	return &unit, nil
}

func GetTerrain(terrainName, ruleset string) (*models.TerrainType, error) {
	var terrain models.TerrainType
	var ok bool

	switch ruleset {
	case "ltt":
		terrain, ok = ltt.TerrainStats[terrainName]
	case "ltx":
		terrain, ok = ltx.TerrainStats[terrainName]
	case "lt75":
		terrain, ok = lt75.TerrainStats[terrainName]
	case "lt76":
		terrain, ok = lt76.TerrainStats[terrainName]
	default:
		return nil, fmt.Errorf("invalid ruleset")
	}

	if !ok {
		return nil, fmt.Errorf("unit \"%s\" not found", terrainName)
	}

	return &terrain, nil
}
