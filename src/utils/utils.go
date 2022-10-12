package utils

import (
	"errors"
	"fmt"

	"freecocoa/src/models"
	"freecocoa/src/stats"
)

func ConvertDefender(data models.DefenderUnit) (*models.UnitDetails, error) {
	answer := models.UnitDetails{}

	unit, ok := stats.UnitStats[data.Name]
	if !ok {
		return nil, fmt.Errorf("unit \"%s\" not found", data.Name)
	}

	if data.VetLevel < 0 || data.VetLevel > 9 {
		return nil, errors.New("veteran level must be an integer between 0 and 9")
	}

	if data.HP < 0 || data.HP > unit.HP {
		return nil, fmt.Errorf("hp must be an integer between 1 and %d (0 or missing defaults to %d)", unit.HP, unit.HP)
	}

	answer.Name = unit.Name
	answer.Attack = unit.Attack * models.VeteranLevels[data.VetLevel]
	answer.Defense = unit.Defense * models.VeteranLevels[data.VetLevel]
	answer.Class = unit.Class
	answer.Class.Name = unit.Class.NameEnum.ToString()
	answer.FP = unit.FP
	answer.CityBuster = unit.CityBuster
	answer.AirAttacker = unit.AirAttacker
	answer.Horse = unit.Horse
	answer.Submarine = unit.Submarine
	answer.BadCityDefender = unit.BadCityDefender
	answer.OnlyNativeAttack = unit.OnlyNativeAttack

	if data.HP == 0 {
		answer.HP = unit.HP
	} else {
		answer.HP = data.HP
	}

	return &answer, nil
}
