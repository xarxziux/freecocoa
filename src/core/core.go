package core

import (
	"freecocoa/src/models"
	"math"
	"math/rand"
)

var veteranLevels = [10]float64{1, 1.5, 1.75, 2, 2.25, 2.5, 2.75, 3, 3.25, 3.5}

func GetStats(av *models.AttackValidated) *models.AttackResults {
	attackResults := models.AttackResults{}

	if av.Attacker.Input.HP == 0 {
		attackResults.Attacker.HP = av.Attacker.Details.HP
	} else {
		attackResults.Attacker.HP = av.Attacker.Input.HP
	}

	if av.Defender.Input.HP == 0 {
		attackResults.Defender.HP = av.Defender.Details.HP
	} else {
		attackResults.Defender.HP = av.Defender.Input.HP
	}

	attackResults.Attacker.FP, attackResults.Defender.FP = setFirepower(av)

	// Apply veteran levels
	attackPower := av.Attacker.Details.AP * veteranLevels[av.Attacker.Input.VetLevel]
	defensePower := av.Defender.Details.DP * veteranLevels[av.Defender.Input.VetLevel]

	// Unit V unit specifics
	if av.Attacker.Details.Horse && av.Defender.Details.Name == "Pikemen" {
		defensePower *= 2
	}

	if av.Attacker.Details.Name == "Submarine" && av.Defender.Details.Name == "Destroyer" {
		defensePower *= 2
	}

	if av.Attacker.Details.AirAttacker && av.Defender.Details.Name == "AEGIS Cruiser" {
		defensePower *= 5
	}

	// Fortification bonuses applies if fortified or in a City
	if av.Defender.Details.Class.CanFortify && (av.Defender.Input.IsFortified || av.Defender.Input.HasCity) {
		defensePower *= 1.5
	}

	// Air, helicopter and missile bypass fortresses
	if av.Defender.Details.Class.CanFortify && av.Defender.Input.HasFortress {
		switch av.Attacker.Details.Class.NameEnum {
		case models.Air, models.Helicopter, models.Missile:
		default:
			defensePower *= 2
		}
	}

	// Airbases only protect against air, helicopter and missile
	if av.Defender.Details.Class.CanFortify &&
		av.Defender.Input.HasAirbase {
		switch av.Attacker.Details.Class.NameEnum {
		case models.Air, models.Helicopter, models.Missile:
			defensePower *= 2
		}
	}

	// Terrain bonuses
	if av.Defender.Details.Class.TerrainDefense {
		defensePower *= getTerrainBonus(av)
	}

	// Get city defense bonus
	if av.Defender.Input.HasCity {
		defensePower *= getCityDefenseBonus(av)
	}

	// Apply movement point penalty
	if av.Attacker.Input.MP < 9 {
		attackPower *= float64(av.Attacker.Input.MP) / 9.0
	}

	attackResults.Attacker.AP = attackPower
	attackResults.Defender.DP = defensePower

	return &attackResults
}

// Based on https://github.com/longturn/freeciv21/blob/51fea1b5f0cedc9361709a574c9d062b6f6f7f7c/common/combat.cpp#L352
func setFirepower(av *models.AttackValidated) (int, int) {
	attFP := av.Attacker.Details.FP
	defFP := av.Defender.Details.FP

	if av.Defender.Input.HasCity {
		if av.Attacker.Details.CityBuster {
			attFP *= 2
		}

		if av.Defender.Details.BadCityDefender {
			attFP *= 2
			defFP = 1
		}
	}

	if av.Attacker.Details.Name == "Fighter" &&
		av.Defender.Details.Name == "Helicopter" {
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

func getCityDefenseBonus(av *models.AttackValidated) float64 {
	citySize := getCitySize(av.City.Size)

	if citySize == models.NoCity {
		panic("Validation error: city size 0")
	}

	bonus := float64(1)

	// Due to the way the logic in effects.ruleset is organised,
	// some of these switch statements filter exceptions and
	// default to on for everything else.
	switch citySize {
	case models.Town:
		switch av.Defender.Details.Class.NameEnum {
		case models.Air, models.DeepSea, models.Helicopter, models.Missile, models.Nuclear, models.Patrol, models.Sea, models.SmallSea, models.Trireme:
		default:
			bonus += 0.5
		}
	case models.City:
		switch av.Defender.Details.Class.NameEnum {
		case models.Air, models.Helicopter, models.Missile, models.Nuclear:
		case models.DeepSea, models.Patrol, models.Sea, models.SmallSea, models.Trireme:
			bonus += 0.5
		default:
			bonus += 1.0
		}
	}

	if av.City.HasWalls {
		switch av.Attacker.Details.Class.NameEnum {
		// TODO: Should DeepSea, Patrol and SmallSea be in here?
		case models.Air, models.Helicopter, models.Missile, models.Sea, models.Trireme:
		default:
			bonus += 1.0
		}
	}

	if av.City.HasCoastalDefense {
		switch av.Attacker.Details.Class.NameEnum {
		// TODO: Should DeepSea, Patrol and SmallSea be in here?
		case models.Sea, models.Trireme:
			bonus += 1.0
		}
	}

	if av.City.HasSAM {
		switch av.Defender.Details.Class.NameEnum {
		case models.Air, models.Helicopter:
			bonus += 1.0
		}
	}

	if av.City.SDILevel > 0 && av.Defender.Details.Class.NameEnum == models.Missile {
		if av.City.IsCapital {
			bonus += 0.7
		} else {
			switch av.City.SDILevel {
			case 1:
				bonus += 0.3
			case 2:
				bonus += 0.6
			case 3:
				bonus += 1.0
			}
		}
	}

	if av.Defender.Input.HasGreatWall {
		switch av.Defender.Details.Class.NameEnum {
		// TODO: Should DeepSea, Patrol and SmallSea be in here?
		case models.Air, models.Helicopter, models.Missile, models.Sea, models.Trireme:
			_ = true
		default:
			bonus += 0.4
		}
	}

	return bonus
}

func getTerrainBonus(av *models.AttackValidated) float64 {
	bonus := (100.0 + float64(av.Terrain.Details.DefenseBonus)) / 100.0

	if av.Terrain.Input.HasRiver {
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

func unitVunit(attacker models.AttackerPair, defender models.DefenderPair) (*models.AttackerPair, *models.DefenderPair) {
	attAP := attacker.Base.AP
	attInitHP := attacker.Base.HP
	attHP := attacker.Base.HP
	attFP := attacker.Base.FP
	attMP := attacker.Base.MP
	defDP := defender.Base.DP
	defInitHP := defender.Base.HP
	defHP := defender.Base.HP
	defFP := defender.Base.FP
	defMP := defender.Base.MP
	pool := attAP + defDP

	for {
		r := rand.Float64() * pool

		if r <= defDP {
			attHP -= defFP
		} else {
			defHP -= attFP
		}

		if attHP <= 0 || defHP <= 0 {
			break
		}
	}

	if attHP <= 0 {
		hpDelta := defInitHP - defHP
		mpDelta := int(math.Round((float64(hpDelta) / float64(defender.Details.HP)) * float64(defender.Details.MP)))

		if mpDelta > defMP {
			mpDelta = defMP
		}

		return nil, &models.DefenderPair{
			Base: models.DefenderBase{
				DP: defDP,
				HP: defHP,
				FP: defFP,
				MP: defMP - mpDelta,
			},
			Details: defender.Details,
		}
	}

	hpDelta := attInitHP - attHP
	mpDelta := int(math.Round((float64(hpDelta) / float64(attacker.Details.HP)) * float64(attacker.Details.MP)))

	if mpDelta > attMP {
		mpDelta = attMP
	}

	return &models.AttackerPair{
		Base: models.AttackerBase{
			AP: attAP,
			HP: attHP,
			FP: attFP,
			MP: attMP - mpDelta,
		},
		Details: attacker.Details,
	}, nil
}
