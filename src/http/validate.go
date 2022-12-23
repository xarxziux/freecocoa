package http

import (
	"errors"
	"fmt"

	"freecocoa/src/models"
	"freecocoa/src/rulesets"
)

func validateInput(avd models.AttackerVDefender) (*models.BaseStats, error) {
	attUnit, ok := rulesets.UnitStats[avd.Attacker.Name]
	if !ok {
		return nil, fmt.Errorf("unit \"%s\" not found", avd.Attacker.Name)
	}

	defUnit, ok := rulesets.UnitStats[avd.Defender.Name]
	if !ok {
		return nil, fmt.Errorf("unit \"%s\" not found", avd.Attacker.Name)
	}

	if avd.Defender.Terrain.Type == "" {
		return nil, fmt.Errorf("no terrain details found")
	}

	terrain, ok := rulesets.TerrainStats[avd.Defender.Terrain.Type]
	if !ok {
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
