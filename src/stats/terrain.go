package stats

// AUTO-GENERATED FILE - DO NOT EDIT
// USE THE SCRIPT parse_terrain_ruleset.js TO RE-GENERATE IF NEEDED

import (
	"github.com/xarxziux/freecocoa/src/models"
)

// TerrainStats lists all units combat stats as extracted from units.ruleset.
var TerrainStats = map[string]models.TerrainType{
	"Inaccessible": {
		Class:        models.LandClass,
		DefenseBonus: 0,
		NoCities:     true,
		UnsafeCoast:  true,
	},
	"Lake": {
		Class:        models.OceanicClass,
		DefenseBonus: 10,
		NoCities:     true,
		NoFortify:    true,
	},
	"Ocean": {
		Class:        models.OceanicClass,
		DefenseBonus: 10,
		NoCities:     true,
		UnsafeCoast:  true,
		NoFortify:    true,
	},
	"Deep Ocean": {
		Class:        models.OceanicClass,
		DefenseBonus: 0,
		NoCities:     true,
		UnsafeCoast:  true,
		NoFortify:    true,
	},
	"Glacier": {
		Class:        models.LandClass,
		DefenseBonus: 0,
		CanHaveRiver: true,
	},
	"Desert": {
		Class:        models.LandClass,
		DefenseBonus: 0,
		CanHaveRiver: true,
	},
	"Forest": {
		Class:        models.LandClass,
		DefenseBonus: 25,
		CanHaveRiver: true,
	},
	"Grassland": {
		Class:        models.LandClass,
		DefenseBonus: 0,
		CanHaveRiver: true,
	},
	"Hills": {
		Class:        models.LandClass,
		DefenseBonus: 50,
		CanHaveRiver: true,
	},
	"Jungle": {
		Class:        models.LandClass,
		DefenseBonus: 25,
		CanHaveRiver: true,
	},
	"Mountains": {
		Class:        models.LandClass,
		DefenseBonus: 100,
		NoCities:     true,
		CanHaveRiver: true,
	},
	"Plains": {
		Class:        models.LandClass,
		DefenseBonus: 0,
		CanHaveRiver: true,
	},
	"Swamp": {
		Class:        models.LandClass,
		DefenseBonus: 25,
		CanHaveRiver: true,
	},
	"Tundra": {
		Class:        models.LandClass,
		DefenseBonus: 0,
		CanHaveRiver: true,
	},
}
