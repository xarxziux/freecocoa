package rulesets

import (
	"errors"
	"fmt"

	"freecocoa/src/models"
)

func PopulateInput(avd models.AttackerVDefender, ruleset string) (*models.BaseStats, error) {
	if avd.Defender.Terrain.Type == "" {
		return nil, fmt.Errorf("no terrain details found")
	}

	attUnit, err := getUnit(avd.Attacker.Name, ruleset)
	if err != nil {
		return nil, err
	}

	defUnit, err := getUnit(avd.Defender.Name, ruleset)
	if err != nil {
		return nil, err
	}

	terrain, err := getTerrain(avd.Defender.Terrain.Type, ruleset)
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

	if hasOnlyOneStructure(avd.Defender.IsFortified, avd.Defender.HasFortress, avd.Defender.HasAirbase, avd.Defender.HasCity) {
		return nil, fmt.Errorf("at most, only one of isFortified, hasFortress, hasAirbase and hasCity can be true")
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

func hasOnlyOneStructure(isFortified, hasFortress, hasAirbase, hasCity bool) bool {
	x := 0

	if isFortified {
		x++
	}

	if hasFortress {
		x++
	}

	if hasAirbase {
		x++
	}

	if hasCity {
		x++
	}

	return x < 2
}
