package lt75

// AUTO-GENERATED FILE - DO NOT EDIT
// USE THE SCRIPT parse_terrain_ruleset.js TO RE-GENERATE IF NEEDED

import (
	"freecocoa/src/models"
)

// TerrainStats lists all units combat stats as extracted from units.ruleset.
var TerrainStats = map[string]models.TerrainType{
	"inaccesible": {
		Name:         "Inaccessible",
		Class:        models.LandClass,
		DefenseBonus: 0,
		NoCities:     true,
		UnsafeCoast:  true,
	},
	"lake": {
		Name:         "Lake",
		Class:        models.OceanicClass,
		DefenseBonus: 10,
		NoCities:     true,
		NoFortify:    true,
	},
	"ocean": {
		Name:         "Ocean",
		Class:        models.OceanicClass,
		DefenseBonus: 10,
		NoCities:     true,
		UnsafeCoast:  true,
		NoFortify:    true,
	},
	"deep_ocean": {
		Name:         "Deep Ocean",
		Class:        models.OceanicClass,
		DefenseBonus: 0,
		NoCities:     true,
		UnsafeCoast:  true,
		NoFortify:    true,
	},
	"glacier": {
		Name:         "Glacier",
		Class:        models.LandClass,
		DefenseBonus: 0,
		CanHaveRiver: true,
	},
	"desert": {
		Name:         "Desert",
		Class:        models.LandClass,
		DefenseBonus: 0,
		CanHaveRiver: true,
	},
	"forest": {
		Name:         "Forest",
		Class:        models.LandClass,
		DefenseBonus: 25,
		CanHaveRiver: true,
	},
	"grassland": {
		Name:         "Grassland",
		Class:        models.LandClass,
		DefenseBonus: 0,
		CanHaveRiver: true,
	},
	"hills": {
		Name:         "Hills",
		Class:        models.LandClass,
		DefenseBonus: 50,
		CanHaveRiver: true,
	},
	"jungle": {
		Name:         "Jungle",
		Class:        models.LandClass,
		DefenseBonus: 25,
		CanHaveRiver: true,
	},
	"mountains": {
		Name:         "Mountains",
		Class:        models.LandClass,
		DefenseBonus: 100,
		NoCities:     true,
		CanHaveRiver: true,
	},
	"plains": {
		Name:         "Plains",
		Class:        models.LandClass,
		DefenseBonus: 0,
		CanHaveRiver: true,
	},
	"swamp": {
		Name:         "Swamp",
		Class:        models.LandClass,
		DefenseBonus: 25,
		CanHaveRiver: true,
	},
	"tundra": {
		Name:         "Tundra",
		Class:        models.LandClass,
		DefenseBonus: 0,
		CanHaveRiver: true,
	},
}
