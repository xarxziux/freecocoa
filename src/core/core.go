package core

import (
	"freecocoa/src/models"
)

var veteranLevels = [10]float64{1, 1.5, 1.75, 2, 2.25, 2.5, 2.75, 3, 3.25, 3.5}

func GetStats(avd models.BaseStats) (*models.FinalStats, error) {
	finalStats := models.FinalStats{}

	finalStats.Attacker.FP, finalStats.Defender.FP = setFirepower(avd)

	if avd.Input.Attacker.HP == 0 {
		finalStats.Attacker.HP = avd.Details.Attacker.HP
	} else {
		finalStats.Attacker.HP = avd.Input.Attacker.HP
	}

	if avd.Input.Defender.HP == 0 {
		finalStats.Defender.HP = avd.Details.Defender.HP
	} else {
		finalStats.Defender.HP = avd.Input.Defender.HP
	}

	// Apply veteran levels
	attackPower := avd.Details.Attacker.AP * veteranLevels[avd.Input.Attacker.VetLevel]
	defensePower := avd.Details.Defender.DP * veteranLevels[avd.Input.Defender.VetLevel]

	// Unit V unit specifics
	if avd.Details.Attacker.Horse && avd.Details.Defender.Name == "Pikemen" {
		defensePower *= 2
	}

	if avd.Details.Attacker.Name == "Submarine" && avd.Details.Defender.Name == "Destroyer" {
		defensePower *= 2
	}

	if avd.Details.Attacker.AirAttacker && avd.Details.Defender.Name == "AEGIS Cruiser" {
		defensePower *= 5
	}

	// Fortification bonuses
	if avd.Details.Defender.Class.CanFortify && avd.Input.Defender.IsFortified {
		defensePower *= 1.5
	}

	if avd.Details.Defender.Class.CanFortify && avd.Input.Defender.HasFortress {
		defensePower *= 2
	}

	// Terrain bonuses
	if avd.Details.Defender.Class.TerrainDefense {
		defensePower *= getTerrainBonus(avd)
	}

	if avd.Input.Defender.HasCity {
		defensePower *= getCityDefenseBonus(avd)
	}

	// Apply movement point penalty
	attackPower *= float64(avd.Input.Attacker.MP) / 9.0

	finalStats.Attacker.AP = attackPower
	finalStats.Defender.DP = defensePower

	return &finalStats, nil
}

// Based on https://github.com/longturn/freeciv21/blob/51fea1b5f0cedc9361709a574c9d062b6f6f7f7c/common/combat.cpp#L352
func setFirepower(base models.BaseStats) (int, int) {
	attFP := base.Details.Attacker.FP
	defFP := base.Details.Defender.FP

	if base.Input.Defender.HasCity {
		if base.Details.Attacker.CityBuster {
			attFP *= 2
		}

		if base.Details.Defender.BadCityDefender {
			attFP *= 2
			defFP = 1
		}
	}

	if base.Details.Attacker.Name == "Fighter" &&
		base.Details.Defender.Name == "Helicopter" {
		defFP = 1
	}

	return attFP, defFP
}

func getCitySize(size int) models.CityType {
	if size == 0 {
		return models.NoCity
	}

	if size < 9 {
		return models.Town
	}

	return models.City
}

func getCityDefenseBonus(avd models.BaseStats) float64 {
	citySize := getCitySize(avd.Input.Defender.City.Size)

	if citySize == models.NoCity {
		panic("Validation error: city size 0")
	}

	bonus := float64(1)

	// Due to the way the logic in effects.ruleset is organised,
	// some of these switch statements filter exceptions and
	// default to on for everything else.
	switch citySize {
	case models.Town:
		switch avd.Details.Defender.Class.NameEnum {
		case models.Air, models.DeepSea, models.Helicopter, models.Missile, models.Nuclear, models.Patrol, models.Sea, models.SmallSea, models.Trireme:
			bonus = 1.0
		default:
			bonus = 1.5
		}
	case models.City:
		switch avd.Details.Defender.Class.NameEnum {
		case models.Air, models.Helicopter, models.Missile:
			bonus = 1.0
		case models.Sea, models.Trireme:
			bonus = 1.5
		default:
			bonus = 2.0
		}
	}

	if avd.Input.Defender.City.HasWalls {
		switch avd.Details.Defender.Class.NameEnum {
		case models.Air, models.Helicopter, models.Missile, models.Sea, models.Trireme:
		default:
			bonus *= 2
		}
	}

	if avd.Input.Defender.City.HasCoastalDefence {
		switch avd.Details.Defender.Class.NameEnum {
		case models.Sea, models.Trireme:
			bonus *= 2
		}
	}

	if avd.Input.Defender.City.HasSAM {
		switch avd.Details.Defender.Class.NameEnum {
		case models.Air, models.Helicopter:
			bonus *= 2
		}
	}

	if avd.Input.Defender.City.SDILevel > 0 && avd.Details.Defender.Class.NameEnum == models.Missile {
		if avd.Input.Defender.City.IsCapital {
			bonus *= 2
		} else {
			switch avd.Input.Defender.City.SDILevel {
			case 1:
				bonus *= 1.3
			case 2:
				bonus *= 1.7
			case 3:
				bonus *= 2.0
			}
		}
	}

	if avd.Input.Defender.HasGreatWall {
		switch avd.Details.Defender.Class.NameEnum {
		case models.Air, models.DeepSea, models.Helicopter, models.Missile, models.Nuclear, models.Patrol, models.Sea, models.SmallSea, models.Trireme:
			_ = true
		default:
			bonus *= 1.4
		}
	}

	return bonus
}

func getTerrainBonus(avd models.BaseStats) float64 {
	bonus := (100.0 + float64(avd.Terrain.DefenseBonus)) / 100.0

	if avd.Input.Defender.Terrain.HasRiver {
		bonus *= 1.5
	}

	return bonus
}
