package utils

import (
	"errors"
	"fmt"

	"freecocoa/src/core"
	"freecocoa/src/models"
	"freecocoa/src/stats"
)

func GetStats(avd models.AttackerVDefender) (*models.AllStats, error) {
	answer := models.AllStats{}

	attUnit, ok := stats.UnitStats[avd.Attacker.Name]
	if !ok {
		return nil, fmt.Errorf("unit \"%s\" not found", avd.Attacker.Name)
	}

	defUnit, ok := stats.UnitStats[avd.Defender.Name]
	if !ok {
		return nil, fmt.Errorf("unit \"%s\" not found", avd.Attacker.Name)
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

	citySize := core.GetCitySize(avd.Defender.City.Size)

	if avd.Defender.HasCity && citySize == models.NoCity {
		return nil, fmt.Errorf("city cannot be of size 0")
	}

	if avd.Attacker.HP == 0 {
		answer.Attacker.HP = attUnit.HP
	} else {
		answer.Attacker.HP = avd.Attacker.HP
	}

	if avd.Defender.HP == 0 {
		answer.Defender.HP = defUnit.HP
	} else {
		answer.Defender.HP = avd.Defender.HP
	}

	answer.Attacker.AP = attUnit.AP
	answer.Defender.DP = defUnit.DP
	answer.Attacker.FP = attUnit.FP
	answer.Defender.FP = defUnit.FP

	return &answer, nil
}
