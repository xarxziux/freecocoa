package http

import (
	"errors"
	"fmt"

	"freecocoa/src/models"
	"freecocoa/src/rulesets"
)

func validateInput(avd models.AttackerVDefender, ruleset string) (*models.BaseStats, error) {
	if avd.Defender.Terrain.Type == "" {
		return nil, fmt.Errorf("no terrain details found")
	}

	attUnit, err := rulesets.GetUnit(avd.Attacker.Name, ruleset)
	if err != nil {
		return nil, err
	}

	defUnit, err := rulesets.GetUnit(avd.Defender.Name, ruleset)
	if err != nil {
		return nil, err
	}

	terrain, err := rulesets.GetTerrain(avd.Defender.Terrain.Type, ruleset)
	if err != nil {
		return nil, err
	}

	if avd.Attacker.VetLevel < 0 ||
		avd.Attacker.VetLevel > 9 ||
		avd.Defender.VetLevel < 0 ||
		avd.Defender.VetLevel > 9 {
		return nil, errors.New("veteran level must be an integer between 0 and 9")
	}

	if avd.Attacker.MP < 0 || avd.Attacker.MP > 9 {
		return nil, errors.New("movement points must be an integer between 0 and 9")
	}

	if avd.Attacker.MP == 0 {
		avd.Attacker.MP = 9
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
		return nil, fmt.Errorf("at most, only one of hasCity, hasFortress and isFortified can be true")
	}

	if avd.Defender.HasCity && avd.Defender.City.Size == 0 {
		return nil, fmt.Errorf("city cannot be of size 0")
	}

	return &models.BaseStats{
		Input: avd,
		Details: models.PairDetails{
			Attacker: *attUnit,
			Defender: *defUnit,
		},
		Terrain: *terrain,
	}, nil
}
