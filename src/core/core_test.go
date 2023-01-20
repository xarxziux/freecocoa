package core_test

import (
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"freecocoa/src/core"
	"freecocoa/src/models"
	"freecocoa/src/rulesets"
)

type sampleData struct {
	AttackerIn  models.AttackerUnit
	DefenderIn  models.DefenderUnit
	AttackerOut models.AttackStats
	DefenderOut models.DefenseStats
}

var _ = Describe("Core", func() {
	Describe("Combat samples", func() {
		It("expected output matches actual output", func() {
			for _, sample := range samples {
				avd := &models.AttackerVDefender{
					Attacker: sample.AttackerIn,
					Defender: sample.DefenderIn,
				}

				fullStats, err := rulesets.PopulateInput(avd, "ltt")
				if err != nil {
					Expect(true).To(Equal(false))
				}

				finalStats := core.GetStats(fullStats)

				Expect(math.Round(finalStats.Attacker.AP*10000) / 10000).To(Equal(sample.AttackerOut.AP))
				Expect(finalStats.Attacker.HP).To(Equal(sample.AttackerOut.HP))
				Expect(finalStats.Attacker.FP).To(Equal(sample.AttackerOut.FP))
				Expect(math.Round(finalStats.Defender.DP*10000) / 10000).To(Equal(sample.DefenderOut.DP))
				Expect(finalStats.Defender.HP).To(Equal(sample.DefenderOut.HP))
				Expect(finalStats.Defender.FP).To(Equal(sample.DefenderOut.FP))
			}
		})
	})
})

var samples = []sampleData{
	{
		// Warriors V warriors
		AttackerIn: models.AttackerUnit{
			Name: "warriors",
		},
		DefenderIn: models.DefenderUnit{
			Name: "warriors",
			Terrain: models.DefenderTerrain{
				Type: "grassland",
			},
		},
		AttackerOut: models.AttackStats{
			AP: 1,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenseStats{
			DP: 1,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Vet bonus
		AttackerIn: models.AttackerUnit{
			Name:     "warriors",
			VetLevel: 3,
		},
		DefenderIn: models.DefenderUnit{
			Name: "warriors",
			Terrain: models.DefenderTerrain{
				Type: "grassland",
			},
		},
		AttackerOut: models.AttackStats{
			AP: 2,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenseStats{
			DP: 1,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Movement penalty
		AttackerIn: models.AttackerUnit{
			Name: "warriors",
			MP:   3,
		},
		DefenderIn: models.DefenderUnit{
			Name: "warriors",
			Terrain: models.DefenderTerrain{
				Type: "grassland",
			},
		},
		AttackerOut: models.AttackStats{
			AP: 0.3333,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenseStats{
			DP: 1,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Terrain bonus - forest
		AttackerIn: models.AttackerUnit{
			Name: "warriors",
		},
		DefenderIn: models.DefenderUnit{
			Name: "warriors",
			Terrain: models.DefenderTerrain{
				Type: "forest",
			},
		},
		AttackerOut: models.AttackStats{
			AP: 1,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenseStats{
			DP: 1.25,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Terrain bonus - mountains
		AttackerIn: models.AttackerUnit{
			Name: "warriors",
		},
		DefenderIn: models.DefenderUnit{
			Name: "warriors",
			Terrain: models.DefenderTerrain{
				Type: "mountains",
			},
		},
		AttackerOut: models.AttackStats{
			AP: 1,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenseStats{
			DP: 2,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Terrain bonus - river
		AttackerIn: models.AttackerUnit{
			Name: "warriors",
		},
		DefenderIn: models.DefenderUnit{
			Name: "warriors",
			Terrain: models.DefenderTerrain{
				Type:     "grassland",
				HasRiver: true,
			},
		},
		AttackerOut: models.AttackStats{
			AP: 1,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenseStats{
			DP: 1.25,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Small city on river with hills
		AttackerIn: models.AttackerUnit{
			Name: "warriors",
		},
		DefenderIn: models.DefenderUnit{
			Name:    "warriors",
			HasCity: true,
			City: models.DefenderCity{
				Size: 5,
			},
			Terrain: models.DefenderTerrain{
				Type:     "hills",
				HasRiver: true,
			},
		},
		AttackerOut: models.AttackStats{
			AP: 1,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenseStats{
			DP: 2.8125,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Small city on river with hills attacked by a slightly damaged veteran catapult
		AttackerIn: models.AttackerUnit{
			Name:     "catapult",
			HP:       8,
			VetLevel: 2,
		},
		DefenderIn: models.DefenderUnit{
			Name:    "warriors",
			HasCity: true,
			City: models.DefenderCity{
				Size: 5,
			},
			Terrain: models.DefenderTerrain{
				Type:     "hills",
				HasRiver: true,
			},
		},
		AttackerOut: models.AttackStats{
			AP: 8.75,
			HP: 8,
			FP: 2,
		},
		DefenderOut: models.DefenseStats{
			DP: 2.8125,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Fortress city with hills attacked by hardcore knights
		AttackerIn: models.AttackerUnit{
			Name:     "knights",
			VetLevel: 6,
		},
		DefenderIn: models.DefenderUnit{
			Name:     "pikemen",
			VetLevel: 3,
			HasCity:  true,
			City: models.DefenderCity{
				Size:     10,
				HasWalls: true,
			},
			Terrain: models.DefenderTerrain{
				Type: "hills",
			},
		},
		AttackerOut: models.AttackStats{
			AP: 16.5,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenseStats{
			DP: 54,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Musketeers in a mountain fortress attacked by dragoons
		AttackerIn: models.AttackerUnit{
			Name:     "dragoons",
			VetLevel: 2,
		},
		DefenderIn: models.DefenderUnit{
			Name:        "musketeers",
			VetLevel:    1,
			IsFortified: true,
			HasFortress: true,
			Terrain: models.DefenderTerrain{
				Type: "mountains",
			},
		},
		AttackerOut: models.AttackStats{
			AP: 8.75,
			HP: 20,
			FP: 1,
		},
		DefenderOut: models.DefenseStats{
			DP: 27,
			HP: 20,
			FP: 1,
		},
	},
	{
		// Fortresses work against ships
		AttackerIn: models.AttackerUnit{
			Name: "destroyer",
		},
		DefenderIn: models.DefenderUnit{
			Name:        "musketeers",
			VetLevel:    1,
			IsFortified: true,
			HasFortress: true,
			Terrain: models.DefenderTerrain{
				Type: "mountains",
			},
		},
		AttackerOut: models.AttackStats{
			AP: 4,
			HP: 30,
			FP: 1,
		},
		DefenderOut: models.DefenseStats{
			DP: 27,
			HP: 20,
			FP: 1,
		},
	},
	{
		// ... but not aircraft
		AttackerIn: models.AttackerUnit{
			Name: "fighter",
		},
		DefenderIn: models.DefenderUnit{
			Name:        "musketeers",
			VetLevel:    1,
			IsFortified: true,
			HasFortress: true,
			Terrain: models.DefenderTerrain{
				Type: "mountains",
			},
		},
		AttackerOut: models.AttackStats{
			AP: 4,
			HP: 20,
			FP: 1,
		},
		DefenderOut: models.DefenseStats{
			DP: 13.5,
			HP: 20,
			FP: 1,
		},
	},
	{
		// Cruise missile V AEGIS cruiser is a mismatch
		AttackerIn: models.AttackerUnit{
			Name: "cruise_missile",
		},
		DefenderIn: models.DefenderUnit{
			Name: "aegis_cruiser",
			Terrain: models.DefenderTerrain{
				Type: "ocean",
			},
		},
		AttackerOut: models.AttackStats{
			AP: 10,
			HP: 5,
			FP: 3,
		},
		DefenderOut: models.DefenseStats{
			DP: 40,
			HP: 30,
			FP: 1,
		},
	},
	{
		// Cruise missile V AEGIS cruiser in a city get a huge FP boost
		AttackerIn: models.AttackerUnit{
			Name: "cruise_missile",
		},
		DefenderIn: models.DefenderUnit{
			Name:    "aegis_cruiser",
			HasCity: true,
			City: models.DefenderCity{
				Size: 10,
			},
			Terrain: models.DefenderTerrain{
				Type: "grassland",
			},
		},
		AttackerOut: models.AttackStats{
			AP: 10,
			HP: 5,
			FP: 12,
		},
		DefenderOut: models.DefenseStats{
			DP: 60,
			HP: 30,
			FP: 1,
		},
	},
}
