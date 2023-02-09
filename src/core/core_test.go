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
	AttackerIn  models.AttackerInput
	DefenderIn  models.DefenderInput
	Terrain     models.TerrainInput
	City        models.CityInput
	AttackerOut models.AttackerBase
	DefenderOut models.DefenderBase
}

var _ = Describe("Core", func() {
	Describe("Combat samples", func() {
		It("expected output matches actual output", func() {
			for _, sample := range samples {
				avd := &models.AttackInput{
					Attacker: sample.AttackerIn,
					Defender: sample.DefenderIn,
					City:     sample.City,
					Terrain:  sample.Terrain,
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
		AttackerIn: models.AttackerInput{
			Name: "warriors",
		},
		DefenderIn: models.DefenderInput{
			Name: "warriors",
		},
		Terrain: models.TerrainInput{
			Type: "grassland",
		},
		AttackerOut: models.AttackerBase{
			AP: 1,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenderBase{
			DP: 1,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Vet bonus
		AttackerIn: models.AttackerInput{
			Name:     "warriors",
			VetLevel: 3,
		},
		DefenderIn: models.DefenderInput{
			Name: "warriors",
		},
		Terrain: models.TerrainInput{
			Type: "grassland",
		},
		AttackerOut: models.AttackerBase{
			AP: 2,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenderBase{
			DP: 1,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Movement penalty
		AttackerIn: models.AttackerInput{
			Name: "warriors",
			MP:   3,
		},
		DefenderIn: models.DefenderInput{
			Name: "warriors",
		},
		Terrain: models.TerrainInput{
			Type: "grassland",
		},
		AttackerOut: models.AttackerBase{
			AP: 0.3333,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenderBase{
			DP: 1,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Terrain bonus - forest
		AttackerIn: models.AttackerInput{
			Name: "warriors",
		},
		DefenderIn: models.DefenderInput{
			Name: "warriors",
		},
		Terrain: models.TerrainInput{
			Type: "forest",
		},
		AttackerOut: models.AttackerBase{
			AP: 1,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenderBase{
			DP: 1.25,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Terrain bonus - mountains
		AttackerIn: models.AttackerInput{
			Name: "warriors",
		},
		DefenderIn: models.DefenderInput{
			Name: "warriors",
		},
		Terrain: models.TerrainInput{
			Type: "mountains",
		},
		AttackerOut: models.AttackerBase{
			AP: 1,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenderBase{
			DP: 2,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Terrain bonus - river
		AttackerIn: models.AttackerInput{
			Name: "warriors",
		},
		DefenderIn: models.DefenderInput{
			Name: "warriors",
		},
		Terrain: models.TerrainInput{
			Type:     "grassland",
			HasRiver: true,
		},
		AttackerOut: models.AttackerBase{
			AP: 1,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenderBase{
			DP: 1.25,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Small city on river with hills
		AttackerIn: models.AttackerInput{
			Name: "warriors",
		},
		DefenderIn: models.DefenderInput{
			Name:    "warriors",
			HasCity: true,
		},
		City: models.CityInput{
			Size: 5,
		},
		Terrain: models.TerrainInput{
			Type:     "hills",
			HasRiver: true,
		},
		AttackerOut: models.AttackerBase{
			AP: 1,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenderBase{
			DP: 2.8125,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Small city on river with hills attacked by a slightly damaged veteran catapult
		AttackerIn: models.AttackerInput{
			Name:     "catapult",
			HP:       8,
			VetLevel: 2,
		},
		DefenderIn: models.DefenderInput{
			Name:    "warriors",
			HasCity: true,
		},
		City: models.CityInput{
			Size: 5,
		},
		Terrain: models.TerrainInput{
			Type:     "hills",
			HasRiver: true,
		},
		AttackerOut: models.AttackerBase{
			AP: 8.75,
			HP: 8,
			FP: 2,
		},
		DefenderOut: models.DefenderBase{
			DP: 2.8125,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Fortress city with hills attacked by hardcore knights
		AttackerIn: models.AttackerInput{
			Name:     "knights",
			VetLevel: 6,
		},
		DefenderIn: models.DefenderInput{
			Name:     "pikemen",
			VetLevel: 3,
			HasCity:  true,
		},
		City: models.CityInput{
			Size:     10,
			HasWalls: true,
		},
		Terrain: models.TerrainInput{
			Type: "hills",
		},
		AttackerOut: models.AttackerBase{
			AP: 16.5,
			HP: 10,
			FP: 1,
		},
		DefenderOut: models.DefenderBase{
			DP: 54,
			HP: 10,
			FP: 1,
		},
	},
	{
		// Musketeers in a mountain fortress attacked by dragoons
		AttackerIn: models.AttackerInput{
			Name:     "dragoons",
			VetLevel: 2,
		},
		DefenderIn: models.DefenderInput{
			Name:        "musketeers",
			VetLevel:    1,
			IsFortified: true,
			HasFortress: true,
		},
		Terrain: models.TerrainInput{
			Type: "mountains",
		},
		AttackerOut: models.AttackerBase{
			AP: 8.75,
			HP: 20,
			FP: 1,
		},
		DefenderOut: models.DefenderBase{
			DP: 27,
			HP: 20,
			FP: 1,
		},
	},
	{
		// Fortresses work against ships
		AttackerIn: models.AttackerInput{
			Name: "destroyer",
		},
		DefenderIn: models.DefenderInput{
			Name:        "musketeers",
			VetLevel:    1,
			IsFortified: true,
			HasFortress: true,
		},
		Terrain: models.TerrainInput{
			Type: "mountains",
		},
		AttackerOut: models.AttackerBase{
			AP: 4,
			HP: 30,
			FP: 1,
		},
		DefenderOut: models.DefenderBase{
			DP: 27,
			HP: 20,
			FP: 1,
		},
	},
	{
		// ... but not aircraft
		AttackerIn: models.AttackerInput{
			Name: "fighter",
		},
		DefenderIn: models.DefenderInput{
			Name:        "musketeers",
			VetLevel:    1,
			IsFortified: true,
			HasFortress: true,
		},
		Terrain: models.TerrainInput{
			Type: "mountains",
		},
		AttackerOut: models.AttackerBase{
			AP: 4,
			HP: 20,
			FP: 1,
		},
		DefenderOut: models.DefenderBase{
			DP: 13.5,
			HP: 20,
			FP: 1,
		},
	},
	{
		// Cruise missile V AEGIS cruiser is a mismatch
		AttackerIn: models.AttackerInput{
			Name: "cruise_missile",
		},
		DefenderIn: models.DefenderInput{
			Name: "aegis_cruiser",
		},
		Terrain: models.TerrainInput{
			Type: "ocean",
		},
		AttackerOut: models.AttackerBase{
			AP: 10,
			HP: 5,
			FP: 3,
		},
		DefenderOut: models.DefenderBase{
			DP: 40,
			HP: 30,
			FP: 1,
		},
	},
	{
		// Cruise missile V AEGIS cruiser in a city get a huge FP boost
		AttackerIn: models.AttackerInput{
			Name: "cruise_missile",
		},
		DefenderIn: models.DefenderInput{
			Name:    "aegis_cruiser",
			HasCity: true,
		},
		City: models.CityInput{
			Size: 10,
		},
		Terrain: models.TerrainInput{
			Type: "grassland",
		},
		AttackerOut: models.AttackerBase{
			AP: 10,
			HP: 5,
			FP: 12,
		},
		DefenderOut: models.DefenderBase{
			DP: 60,
			HP: 30,
			FP: 1,
		},
	},
}
