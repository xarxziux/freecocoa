package core

import (
	"freecocoa/src/models"
	"math"
)

var veteranLevels = [10]float64{1, 1.5, 1.75, 2, 2.25, 2.5, 2.75, 3, 3.25, 3.5}

func GetStats(avd *models.BaseStats) *models.FinalStats {
	finalStats := models.FinalStats{}

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
	finalStats.Attacker.FP, finalStats.Defender.FP = setFirepower(avd)

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

	// Fortification bonuses applies if fortified or in a City
	if avd.Details.Defender.Class.CanFortify && (avd.Input.Defender.IsFortified || avd.Input.Defender.HasCity) {
		defensePower *= 1.5
	}

	// Air, helicopter and missile bypass fortresses
	if avd.Details.Defender.Class.CanFortify && avd.Input.Defender.HasFortress {
		switch avd.Details.Attacker.Class.NameEnum {
		case models.Air, models.Helicopter, models.Missile:
		default:
			defensePower *= 2
		}
	}

	// Airbases only protect against air, helicopter and missile
	if avd.Details.Defender.Class.CanFortify &&
		avd.Input.Defender.HasAirbase {
		switch avd.Details.Attacker.Class.NameEnum {
		case models.Air, models.Helicopter, models.Missile:
			defensePower *= 2
		}
	}

	// Terrain bonuses
	if avd.Details.Defender.Class.TerrainDefense {
		defensePower *= getTerrainBonus(avd)
	}

	// Get city defense bonus
	if avd.Input.Defender.HasCity {
		defensePower *= getCityDefenseBonus(avd)
	}

	// Apply movement point penalty
	attackPower *= float64(avd.Input.Attacker.MP) / 9.0

	finalStats.Attacker.AP = attackPower
	finalStats.Defender.DP = defensePower

	return &finalStats
}

// Based on https://github.com/longturn/freeciv21/blob/51fea1b5f0cedc9361709a574c9d062b6f6f7f7c/common/combat.cpp#L352
func setFirepower(base *models.BaseStats) (int, int) {
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

func getCityDefenseBonus(avd *models.BaseStats) float64 {
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
		default:
			bonus += 0.5
		}
	case models.City:
		switch avd.Details.Defender.Class.NameEnum {
		case models.Air, models.Helicopter, models.Missile, models.Nuclear:
		case models.DeepSea, models.Patrol, models.Sea, models.SmallSea, models.Trireme:
			bonus += 0.5
		default:
			bonus += 1.0
		}
	}

	if avd.Input.Defender.City.HasWalls {
		switch avd.Details.Attacker.Class.NameEnum {
		// TODO: Should DeepSea, Patrol and SmallSea be in here?
		case models.Air, models.Helicopter, models.Missile, models.Sea, models.Trireme:
		default:
			bonus += 1.0
		}
	}

	if avd.Input.Defender.City.HasCoastalDefense {
		switch avd.Details.Attacker.Class.NameEnum {
		// TODO: Should DeepSea, Patrol and SmallSea be in here?
		case models.Sea, models.Trireme:
			bonus += 1.0
		}
	}

	if avd.Input.Defender.City.HasSAM {
		switch avd.Details.Defender.Class.NameEnum {
		case models.Air, models.Helicopter:
			bonus += 1.0
		}
	}

	if avd.Input.Defender.City.SDILevel > 0 && avd.Details.Defender.Class.NameEnum == models.Missile {
		if avd.Input.Defender.City.IsCapital {
			bonus += 0.7
		} else {
			switch avd.Input.Defender.City.SDILevel {
			case 1:
				bonus += 0.3
			case 2:
				bonus += 0.6
			case 3:
				bonus += 1.0
			}
		}
	}

	if avd.Input.Defender.HasGreatWall {
		switch avd.Details.Defender.Class.NameEnum {
		// TODO: Should DeepSea, Patrol and SmallSea be in here?
		case models.Air, models.Helicopter, models.Missile, models.Sea, models.Trireme:
			_ = true
		default:
			bonus += 0.4
		}
	}

	return bonus
}

func getTerrainBonus(avd *models.BaseStats) float64 {
	bonus := (100.0 + float64(avd.Terrain.DefenseBonus)) / 100.0

	if avd.Input.Defender.Terrain.HasRiver {
		bonus *= 1.25
	}

	return bonus
}

func GetBuildCosts(bci *models.BuildCostInput) *models.BuildCostOutput {
	bco := models.BuildCostOutput{}
	cost := bci.ShieldCost
	inc := bci.ShieldOutput
	current := bci.ShieldCurrent
	remaining := bci.ShieldCost - current
	bcItems := make([]models.BuildCostItem, 0, 10)

	if current >= cost {
		bcItems = append(bcItems,
			models.BuildCostItem{
				ShieldCurrent:   current,
				ShieldRemaining: 0,
				BuyCost:         0,
			},
		)

		bco.BuildCosts = bcItems

		return &bco
	}

	bcItems = append(bcItems,
		models.BuildCostItem{
			ShieldCurrent:   current,
			ShieldRemaining: remaining,
			BuyCost:         getBuyCost(remaining),
		},
	)

	for i := 1; i < 10; i++ {
		current += inc
		remaining -= inc

		if current >= cost {
			bcItems = append(bcItems,
				models.BuildCostItem{
					ShieldCurrent:   current,
					ShieldRemaining: 0,
					BuyCost:         0,
				},
			)

			break
		}

		bcItems = append(bcItems,
			models.BuildCostItem{
				ShieldCurrent:   current,
				ShieldRemaining: remaining,
				BuyCost:         getBuyCost(remaining),
			},
		)
	}

	bco.BuildCosts = bcItems

	return &bco
}

func getBuyCost(shields int) int {
	return int(math.Round((2.0 * float64(shields)) + (math.Pow(float64(shields), 2) / 20.0)))
}
