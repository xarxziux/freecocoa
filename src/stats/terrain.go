package stats

import (
	"github.com/xarxziux/freecoco/src/models"
)

// TerrainStats lists all units combat stats as extracted from units.ruleset.
var TerrainStats = map[string]models.TerrainType{
	"inaccesible": {
		Name:         "Inaccessible",
		Class:        models.LandClass,
		DefenseBonus: 0,
	},
	"lake": {
		Name:         "Lake",
		Class:        models.OceanicClass,
		DefenseBonus: 10,
	},
	"ocean": {
		Name:         "Ocean",
		Class:        models.OceanicClass,
		DefenseBonus: 10,
	},
	"deep_ocean": {
		Name:         "Deep Ocean",
		Class:        models.OceanicClass,
		DefenseBonus: 0,
	},
	"glacier": {
		Name:         "Glacier",
		Class:        models.LandClass,
		DefenseBonus: 0,
	},
	"desert": {
		Name:         "Desert",
		Class:        models.LandClass,
		DefenseBonus: 0,
	},
	"forest": {
		Name:         "Forest",
		Class:        models.LandClass,
		DefenseBonus: 25,
	},
	"grassland": {
		Name:         "Grassland",
		Class:        models.LandClass,
		DefenseBonus: 0,
	},
	"hills": {
		Name:         "Hills",
		Class:        models.LandClass,
		DefenseBonus: 50,
	},
	"jungle": {
		Name:         "Jungle",
		Class:        models.LandClass,
		DefenseBonus: 25,
	},
	"mountains": {
		Name:         "Mountains",
		Class:        models.LandClass,
		DefenseBonus: 100,
	},
	"plains": {
		Name:         "Plains",
		Class:        models.LandClass,
		DefenseBonus: 0,
	},
	"swamp": {
		Name:         "Swamp",
		Class:        models.LandClass,
		DefenseBonus: 25,
	},
	"tundra": {
		Name:         "Tundra",
		Class:        models.LandClass,
		DefenseBonus: 0,
	},
}
