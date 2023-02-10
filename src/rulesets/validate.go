package rulesets

import (
	"errors"
	"fmt"

	"freecocoa/src/models"
)

func populateAttacker(attacker models.AttackerInput, ruleset string) (*models.AttackerValidated, error) {
	validated := models.AttackerValidated{}

	attackerDetails, err := getUnit(attacker.Name, ruleset)
	if err != nil {
		return nil, err
	}

	if attacker.VetLevel < 0 || attacker.VetLevel > 9 {
		return nil, errors.New("veteran level must be an integer between 0 and 9")
	}

	attackerMaxMP := attackerDetails.MP + attacker.VetLevel

	if attacker.MP < 0 || attacker.MP > attackerMaxMP {
		return nil, fmt.Errorf("movement points must be between 0 and %d", attackerMaxMP)
	}

	if attacker.MP == 0 {
		attacker.MP = attackerMaxMP
	}

	if attacker.HP < 0 || attacker.HP > attackerDetails.HP {
		return nil, fmt.Errorf("attacker HP must be an integer between 1 and %d (0 or missing defaults to %d)", attackerDetails.HP, attackerDetails.HP)
	}

	validated.Input = attacker
	validated.Details = *attackerDetails

	return &validated, nil
}

func populateDefender(defender models.DefenderInput, ruleset string) (*models.DefenderValidated, error) {
	validated := models.DefenderValidated{}

	defenderDetails, err := getUnit(defender.Name, ruleset)
	if err != nil {
		return nil, err
	}

	if defender.VetLevel < 0 || defender.VetLevel > 9 {
		return nil, errors.New("veteran level must be an integer between 0 and 9")
	}

	if defender.HP < 0 || defender.HP > defenderDetails.HP {
		return nil, fmt.Errorf("defender HP must be an integer between 1 and %d (0 or missing defaults to %d)", defenderDetails.HP, defenderDetails.HP)
	}

	if tooManyStructures(defender.HasFortress, defender.HasAirbase, defender.HasCity) {
		return nil, fmt.Errorf("at most, only one of hasFortress, hasAirbase and hasCity can be true")
	}

	validated.Input = defender
	validated.Details = *defenderDetails
	return &validated, nil
}

func populateTerrain(terrain models.TerrainInput, ruleset string) (*models.TerrainValidated, error) {
	validated := models.TerrainValidated{}

	terrainDetails, err := getTerrain(terrain.Type, ruleset)
	if err != nil {
		return nil, err
	}

	validated.Input = terrain
	validated.Details = *terrainDetails
	return &validated, nil
}

func PopulateInput(input *models.AttackInput, ruleset string) (*models.AttackValidated, error) {
	validated := models.AttackValidated{}

	attacker, err := populateAttacker(input.Attacker, ruleset)
	if err != nil {
		return nil, err
	}

	defender, err := populateDefender(input.Defender, ruleset)
	if err != nil {
		return nil, err
	}

	if input.Terrain.Type == "" {
		return nil, fmt.Errorf("no terrain details found")
	}

	terrain, err := populateTerrain(input.Terrain, ruleset)
	if err != nil {
		return nil, err
	}

	if input.Defender.HasCity && input.City.Size == 0 {
		return nil, fmt.Errorf("city cannot be of size 0")
	}

	// No fortresses on rivers
	// if input.Defender.Terrain.HasRiver && input.Defender.HasFortress {
	//	return nil, errors.New("cannot have a fortress on a river")
	// }

	validated.Attacker = *attacker
	validated.Defender = *defender
	validated.Terrain = *terrain

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

func PopulateSim(input *models.AttackAllInput, ruleset string) (*models.AttackAllValidated, error) {
	validated := models.AttackAllValidated{}
	attackers := make([]*models.AttackerValidated, 0, len(input.Attackers))
	defenders := make([]*models.DefenderValidated, 0, len(input.Defenders))

	for _, unit := range input.Attackers {
		attacker, err := populateAttacker(unit, ruleset)

		if err != nil {
			return nil, err
		}

		attackers = append(attackers, attacker)
	}

	for _, unit := range input.Defenders {
		defender, err := populateDefender(unit, ruleset)

		if err != nil {
			return nil, err
		}

		defenders = append(defenders, defender)
	}

	terrain, err := populateTerrain(input.Terrain, ruleset)
	if err != nil {
		return nil, err
	}

	validated.Attackers = attackers
	validated.Defenders = defenders
	validated.Terrain = *terrain

	if input.City.Size == 0 {
		validated.City = input.City
	} else {
		validated.City = models.CityInput{}
	}

	return nil, nil
}
