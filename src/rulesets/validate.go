package rulesets

import (
	"errors"
	"fmt"

	"freecocoa/src/models"
)

func PopulateInput(input *models.AttackInput, ruleset string) (*models.AttackValidated, error) {
	validated := models.AttackValidated{}

	if input.Terrain.Type == "" {
		return nil, fmt.Errorf("no terrain details found")
	}

	attackerDetails, err := getUnit(input.Attacker.Name, ruleset)
	if err != nil {
		return nil, err
	}

	defenderDetails, err := getUnit(input.Defender.Name, ruleset)
	if err != nil {
		return nil, err
	}

	terrain, err := getTerrain(input.Terrain.Type, ruleset)
	if err != nil {
		return nil, err
	}

	if input.Attacker.VetLevel < 0 ||
		input.Attacker.VetLevel > 9 ||
		input.Defender.VetLevel < 0 ||
		input.Defender.VetLevel > 9 {
		return nil, errors.New("veteran level must be an integer between 0 and 9")
	}

	attackerMaxMP := attackerDetails.MP + input.Attacker.VetLevel

	if input.Attacker.MP < 0 || input.Attacker.MP > attackerMaxMP {
		return nil, fmt.Errorf("movement points must be between 0 and %d", attackerMaxMP)
	}

	if input.Attacker.MP == 0 {
		input.Attacker.MP = attackerMaxMP
	}

	if input.Attacker.HP < 0 || input.Attacker.HP > attackerDetails.HP {
		return nil, fmt.Errorf("attacker HP must be an integer between 1 and %d (0 or missing defaults to %d)", attackerDetails.HP, attackerDetails.HP)
	}

	if input.Defender.HP < 0 || input.Defender.HP > defenderDetails.HP {
		return nil, fmt.Errorf("defender HP must be an integer between 1 and %d (0 or missing defaults to %d)", defenderDetails.HP, defenderDetails.HP)
	}

	if tooManyStructures(input.Defender.HasFortress, input.Defender.HasAirbase, input.Defender.HasCity) {
		return nil, fmt.Errorf("at most, only one of hasFortress, hasAirbase and hasCity can be true")
	}

	if input.Defender.HasCity && input.City.Size == 0 {
		return nil, fmt.Errorf("city cannot be of size 0")
	}

	// No fortresses on rivers
	// if input.Defender.Terrain.HasRiver && input.Defender.HasFortress {
	//	return nil, errors.New("cannot have a fortress on a river")
	// }

	validated.Attacker.Input = input.Attacker
	validated.Attacker.Details = *attackerDetails
	validated.Defender.Input = input.Defender
	validated.Defender.Details = *defenderDetails
	validated.Terrain.Input = input.Terrain
	validated.Terrain.Details = *terrain

	if input.Defender.HasCity {
		validated.City = input.City
	} else {
		validated.City = models.CityInput{}
	}

	return &validated, nil
}

func PopulateBuildCost(bci *models.BuildCostInput, ruleset string) error {
	if bci.ShieldCost < 0 ||
		bci.ShieldOutput < 0 ||
		bci.ShieldCurrent < 0 {
		return errors.New("int values cannot be negative")
	}

	if bci.UnitName == "" {
		if bci.ShieldCost == 0 {
			return errors.New("no name or cost value provided")
		}

		return nil
	}

	unit, err := getUnit(bci.UnitName, ruleset)
	if err != nil {
		return err
	}

	bci.ShieldCost = unit.BuildCost
	return nil
}

func tooManyStructures(hasFortress, hasAirbase, hasCity bool) bool {
	x := 0

	if hasFortress {
		x++
	}

	if hasAirbase {
		x++
	}

	if hasCity {
		x++
	}

	return x > 1
}
