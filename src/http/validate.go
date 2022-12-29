package http

import (
	"errors"
	"fmt"

	"freecocoa/src/models"
	"freecocoa/src/rulesets/lt75"
	"freecocoa/src/rulesets/lt76"
	"freecocoa/src/rulesets/ltt"
	"freecocoa/src/rulesets/ltx"
)

func validateInput(avd models.AttackerVDefender, ruleset string) (*models.BaseStats, error) {
	var attUnit models.UnitDetails
	var defUnit models.UnitDetails
	var terrain models.TerrainType
	var attOK bool
	var defOK bool
	var terrOK bool

	if avd.Defender.Terrain.Type == "" {
		return nil, fmt.Errorf("no terrain details found")
	}

	switch ruleset {
	case "ltt":
		attUnit, attOK = ltt.UnitStats[avd.Attacker.Name]
		defUnit, defOK = ltt.UnitStats[avd.Defender.Name]
		terrain, terrOK = ltt.TerrainStats[avd.Defender.Terrain.Type]
	case "ltx":
		attUnit, attOK = ltx.UnitStats[avd.Attacker.Name]
		defUnit, defOK = ltx.UnitStats[avd.Defender.Name]
		terrain, terrOK = ltx.TerrainStats[avd.Defender.Terrain.Type]
	case "lt75":
		attUnit, attOK = lt75.UnitStats[avd.Attacker.Name]
		defUnit, defOK = lt75.UnitStats[avd.Defender.Name]
		terrain, terrOK = lt75.TerrainStats[avd.Defender.Terrain.Type]
	case "lt76":
		attUnit, attOK = lt76.UnitStats[avd.Attacker.Name]
		defUnit, defOK = lt76.UnitStats[avd.Defender.Name]
		terrain, terrOK = lt76.TerrainStats[avd.Defender.Terrain.Type]
	default:
		panic("validation failure: bad ruleset")
	}

	if !attOK {
		return nil, fmt.Errorf("unit \"%s\" not found", avd.Attacker.Name)
	}

	if !defOK {
		return nil, fmt.Errorf("unit \"%s\" not found", avd.Attacker.Name)
	}

	if !terrOK {
		return nil, fmt.Errorf("unrecognised terrain type \"%s\"", avd.Defender.Terrain.Type)
	}

	if avd.Attacker.VetLevel > 9 || avd.Defender.VetLevel > 9 {
		return nil, errors.New("veteran level must be an integer between 0 and 9")
	}

	if avd.Attacker.HP > attUnit.HP {
		return nil, fmt.Errorf("attacker HP must be an integer between 1 and %d (0 or missing defaults to %d)", attUnit.HP, attUnit.HP)
	}

	if avd.Defender.HP > defUnit.HP {
		return nil, fmt.Errorf("defender HP must be an integer between 1 and %d (0 or missing defaults to %d)", defUnit.HP, defUnit.HP)
	}

	if (avd.Defender.HasCity && avd.Defender.HasFortress) ||
		(avd.Defender.HasCity && avd.Defender.IsFortified) ||
		(avd.Defender.HasFortress && avd.Defender.IsFortified) {
		return nil, fmt.Errorf("only one of HasCity, HasFortress and IsFortified can be true")
	}

	if avd.Defender.HasCity && avd.Defender.City.Size == 0 {
		return nil, fmt.Errorf("city cannot be of size 0")
	}

	return &models.BaseStats{
		Input: avd,
		Details: models.PairDetails{
			Attacker: attUnit,
			Defender: defUnit,
		},
		Terrain: terrain,
	}, nil
}
