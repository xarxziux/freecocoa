package stats

// AUTO-GENERATED FILE - DO NOT EDIT
// USE THE SCRIPT parse_units_ruleset.js TO RE-GENERATE IF NEEDED

import (
	"github.com/xarxziux/freecocoa/src/models"
)

// UnitClassStats lists all unit classes and combat-relevant flags as extracted
// from units.ruleset.
var UnitClassStats = map[string]models.UnitClass{
	"missile": {
		Missile:     true,
		Unreachable: true,
	},
	"land": {
		TerrainDefense: true,
		CanFortify:     true,
	},
	"land_small": {
		TerrainDefense: true,
		CanFortify:     true,
	},
	"land_big": {
		CanFortify: true,
	},
	"land_siege": {
		CanFortify:      true,
		AttackNonNative: true,
	},
	"merchant": {
		TerrainDefense: true,
	},
	"sea": {
		AttackNonNative: true,
	},
	"trireme": {},
	"heli": {
		Unreachable: true,
	},
	"air": {
		Unreachable: true,
	},
}

// UnitStats lists all units combat stats and combat-relevant flags as extracted
// from units.ruleset.
var UnitStats = map[string]models.UnitDetails{
	"settlers": {
		Name:    "Settlers",
		Class:   UnitClassStats["SmallLand"],
		Attack:  0,
		Defense: 1,
		HP:      20,
		FP:      1,
	},
	"migrants": {
		Name:    "Migrants",
		Class:   UnitClassStats["SmallLand"],
		Attack:  0,
		Defense: 1,
		HP:      20,
		FP:      1,
	},
	"immigrants": {
		Name:    "Immigrants",
		Class:   UnitClassStats["SmallLand"],
		Attack:  0,
		Defense: 1,
		HP:      20,
		FP:      1,
	},
	"tribal_worker": {
		Name:    "Tribal Workers",
		Class:   UnitClassStats["SmallLand"],
		Attack:  0,
		Defense: 1,
		HP:      10,
		FP:      1,
	},
	"worker": {
		Name:    "Workers",
		Class:   UnitClassStats["SmallLand"],
		Attack:  0,
		Defense: 1,
		HP:      10,
		FP:      1,
	},
	"engineers": {
		Name:    "Engineers",
		Class:   UnitClassStats["SmallLand"],
		Attack:  0,
		Defense: 2,
		HP:      20,
		FP:      1,
	},
	"warriors": {
		Name:    "Warriors",
		Class:   UnitClassStats["Land"],
		Attack:  1,
		Defense: 1,
		HP:      10,
		FP:      1,
	},
	"phalanx": {
		Name:    "Phalanx",
		Class:   UnitClassStats["Land"],
		Attack:  1,
		Defense: 2,
		HP:      10,
		FP:      1,
	},
	"archers": {
		Name:    "Archers",
		Class:   UnitClassStats["Land"],
		Attack:  3,
		Defense: 1,
		HP:      10,
		FP:      1,
	},
	"legion": {
		Name:    "Swordsmen",
		Class:   UnitClassStats["Land"],
		Attack:  4,
		Defense: 2,
		HP:      10,
		FP:      1,
	},
	"pikemen": {
		Name:    "Pikemen",
		Class:   UnitClassStats["Land"],
		Attack:  2,
		Defense: 3,
		HP:      10,
		FP:      1,
	},
	"musketeers": {
		Name:    "Musketeers",
		Class:   UnitClassStats["Land"],
		Attack:  3,
		Defense: 3,
		HP:      20,
		FP:      1,
	},
	"riflemen": {
		Name:    "Riflemen",
		Class:   UnitClassStats["Land"],
		Attack:  5,
		Defense: 4,
		HP:      20,
		FP:      1,
	},
	"infantry": {
		Name:    "Infantry",
		Class:   UnitClassStats["Land"],
		Attack:  5,
		Defense: 6,
		HP:      20,
		FP:      1,
	},
	"alpine_troops": {
		Name:    "Alpine Troops",
		Class:   UnitClassStats["Land"],
		Attack:  7,
		Defense: 4,
		HP:      20,
		FP:      1,
	},
	"partisan": {
		Name:    "Partisan",
		Class:   UnitClassStats["Land"],
		Attack:  4,
		Defense: 5,
		HP:      20,
		FP:      1,
	},
	"fanatics": {
		Name:    "Fanatics",
		Class:   UnitClassStats["Land"],
		Attack:  3,
		Defense: 4,
		HP:      20,
		FP:      1,
	},
	"marines": {
		Name:    "Marines",
		Class:   UnitClassStats["Land"],
		Attack:  8,
		Defense: 5,
		HP:      20,
		FP:      1,
	},
	"paratroopers": {
		Name:    "Paratroopers",
		Class:   UnitClassStats["Land"],
		Attack:  6,
		Defense: 4,
		HP:      20,
		FP:      1,
	},
	"mech_inf": {
		Name:    "Mech. Inf.",
		Class:   UnitClassStats["BigLand"],
		Attack:  6,
		Defense: 6,
		HP:      30,
		FP:      1,
	},
	"horsemen": {
		Name:    "Horsemen",
		Class:   UnitClassStats["Land"],
		Attack:  2,
		Defense: 1,
		HP:      10,
		FP:      1,
		Horse:   true,
	},
	"chariot": {
		Name:    "Chariot",
		Class:   UnitClassStats["BigLand"],
		Attack:  4,
		Defense: 1,
		HP:      10,
		FP:      1,
		Horse:   true,
	},
	"elephants": {
		Name:    "Elephants",
		Class:   UnitClassStats["Land"],
		Attack:  3,
		Defense: 2,
		HP:      10,
		FP:      1,
		Horse:   true,
	},
	"crusaders": {
		Name:    "Crusaders",
		Class:   UnitClassStats["Land"],
		Attack:  3,
		Defense: 2,
		HP:      15,
		FP:      1,
		Horse:   true,
	},
	"knights": {
		Name:    "Knights",
		Class:   UnitClassStats["Land"],
		Attack:  6,
		Defense: 3,
		HP:      10,
		FP:      1,
		Horse:   true,
	},
	"dragoons": {
		Name:    "Dragoons",
		Class:   UnitClassStats["Land"],
		Attack:  5,
		Defense: 2,
		HP:      20,
		FP:      1,
		Horse:   true,
	},
	"cavalry": {
		Name:    "Cavalry",
		Class:   UnitClassStats["Land"],
		Attack:  8,
		Defense: 3,
		HP:      20,
		FP:      1,
		Horse:   true,
	},
	"armor": {
		Name:    "Armor",
		Class:   UnitClassStats["BigLand"],
		Attack:  10,
		Defense: 5,
		HP:      30,
		FP:      1,
	},
	"catapult": {
		Name:       "Catapult",
		Class:      UnitClassStats["BigSiege"],
		Attack:     5,
		Defense:    1,
		HP:         10,
		FP:         1,
		CityBuster: true,
	},
	"cannon": {
		Name:       "Cannon",
		Class:      UnitClassStats["BigSiege"],
		Attack:     8,
		Defense:    1,
		HP:         20,
		FP:         1,
		CityBuster: true,
	},
	"artillery": {
		Name:       "Artillery",
		Class:      UnitClassStats["BigSiege"],
		Attack:     10,
		Defense:    2,
		HP:         20,
		FP:         1,
		CityBuster: true,
	},
	"howitzer": {
		Name:       "Howitzer",
		Class:      UnitClassStats["BigSiege"],
		Attack:     12,
		Defense:    2,
		HP:         30,
		FP:         1,
		CityBuster: true,
	},
	"fighter": {
		Name:        "Fighter",
		Class:       UnitClassStats["Air"],
		Attack:      4,
		Defense:     4,
		HP:          20,
		FP:          1,
		AirAttacker: true,
	},
	"bomber": {
		Name:        "Bomber",
		Class:       UnitClassStats["Air"],
		Attack:      6,
		Defense:     1,
		HP:          30,
		FP:          1,
		AirAttacker: true,
	},
	"helicopter": {
		Name:        "Helicopter",
		Class:       UnitClassStats["Helicopter"],
		Attack:      5,
		Defense:     3,
		HP:          30,
		FP:          1,
		AirAttacker: true,
	},
	"stealth_fighter": {
		Name:        "Stealth Fighter",
		Class:       UnitClassStats["Air"],
		Attack:      8,
		Defense:     7,
		HP:          20,
		FP:          1,
		AirAttacker: true,
	},
	"stealth_bomber": {
		Name:        "Stealth Bomber",
		Class:       UnitClassStats["Air"],
		Attack:      9,
		Defense:     3,
		HP:          30,
		FP:          1,
		AirAttacker: true,
	},
	"trireme": {
		Name:            "Trireme",
		Class:           UnitClassStats["Trireme"],
		Attack:          1,
		Defense:         1,
		HP:              10,
		FP:              1,
		BadCityDefender: true,
	},
	"barge": {
		Name:            "Barge",
		Class:           UnitClassStats["Trireme"],
		Attack:          0,
		Defense:         2,
		HP:              20,
		FP:              1,
		BadCityDefender: true,
	},
	"longboat": {
		Name:            "Longboat",
		Class:           UnitClassStats["Sea"],
		Attack:          1,
		Defense:         2,
		HP:              4,
		FP:              1,
		BadCityDefender: true,
	},
	"square_rigged_caravel": {
		Class:           UnitClassStats["Sea"],
		Attack:          3,
		Defense:         2,
		HP:              8,
		FP:              1,
		BadCityDefender: true,
	},
	"caravel": {
		Name:            "Caravel",
		Class:           UnitClassStats["Sea"],
		Attack:          1,
		Defense:         2,
		HP:              10,
		FP:              1,
		BadCityDefender: true,
	},
	"galleon": {
		Name:            "Galleon",
		Class:           UnitClassStats["Sea"],
		Attack:          0,
		Defense:         2,
		HP:              20,
		FP:              1,
		BadCityDefender: true,
	},
	"frigate": {
		Name:            "Frigate",
		Class:           UnitClassStats["Sea"],
		Attack:          4,
		Defense:         3,
		HP:              20,
		FP:              1,
		BadCityDefender: true,
	},
	"flagship_frigate": {
		Name:            "Flagship Frigate",
		Class:           UnitClassStats["Sea"],
		Attack:          5,
		Defense:         3,
		HP:              22,
		FP:              1,
		BadCityDefender: true,
	},
	"ironclad": {
		Name:            "Ironclad",
		Class:           UnitClassStats["Sea"],
		Attack:          5,
		Defense:         6,
		HP:              20,
		FP:              1,
		BadCityDefender: true,
	},
	"destroyer": {
		Name:            "Destroyer",
		Class:           UnitClassStats["Sea"],
		Attack:          4,
		Defense:         4,
		HP:              30,
		FP:              1,
		BadCityDefender: true,
	},
	"cruiser": {
		Name:            "Cruiser",
		Class:           UnitClassStats["Sea"],
		Attack:          4,
		Defense:         8,
		HP:              30,
		FP:              1,
		BadCityDefender: true,
	},
	"aegis_cruiser": {
		Name:            "AEGIS Cruiser",
		Class:           UnitClassStats["Sea"],
		Attack:          8,
		Defense:         8,
		HP:              30,
		FP:              1,
		BadCityDefender: true,
	},
	"battleship": {
		Name:            "Battleship",
		Class:           UnitClassStats["Sea"],
		Attack:          14,
		Defense:         14,
		HP:              30,
		FP:              2,
		BadCityDefender: true,
	},
	"submarine": {
		Name:            "Submarine",
		Class:           UnitClassStats["Sea"],
		Attack:          12,
		Defense:         4,
		HP:              20,
		FP:              1,
		Submarine:       true,
		BadCityDefender: true,
	},
	"nuclear_submarine": {
		Name:            "Nuclear Submarine",
		Class:           UnitClassStats["Sea"],
		Attack:          12,
		Defense:         4,
		HP:              20,
		FP:              1,
		Submarine:       true,
		BadCityDefender: true,
	},
	"carrier": {
		Name:            "Carrier",
		Class:           UnitClassStats["Sea"],
		Attack:          0,
		Defense:         10,
		HP:              40,
		FP:              1,
		BadCityDefender: true,
	},
	"transport": {
		Name:            "Transport",
		Class:           UnitClassStats["Sea"],
		Attack:          0,
		Defense:         3,
		HP:              30,
		FP:              1,
		BadCityDefender: true,
	},
	"missile": {
		Name:        "Missile",
		Class:       UnitClassStats["Missile"],
		Attack:      15,
		Defense:     0,
		HP:          2,
		FP:          2,
		AirAttacker: true,
	},
	"cruise_missile": {
		Name:        "Cruise Missile",
		Class:       UnitClassStats["Missile"],
		Attack:      10,
		Defense:     0,
		HP:          5,
		FP:          3,
		CityBuster:  true,
		AirAttacker: true,
	},
	"intercontinental_missile": {
		Name:        "Intercontinental Missile",
		Class:       UnitClassStats["Missile"],
		Attack:      25,
		Defense:     0,
		HP:          5,
		FP:          3,
		AirAttacker: true,
	},
	"nuclear_bomb": {
		Name:    "Nuclear Bomb",
		Class:   UnitClassStats["Missile"],
		Attack:  99,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"nuclear": {
		Name:        "Nuclear",
		Class:       UnitClassStats["Missile"],
		Attack:      99,
		Defense:     0,
		HP:          10,
		FP:          1,
		AirAttacker: true,
	},
	"diplomat": {
		Name:    "Diplomat",
		Class:   UnitClassStats["SmallLand"],
		Attack:  0,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"spy": {
		Name:    "Spy",
		Class:   UnitClassStats["SmallLand"],
		Attack:  0,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"operative": {
		Name:    "Operative",
		Class:   UnitClassStats["SmallLand"],
		Attack:  0,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"scribe": {
		Name:    "Scribe",
		Class:   UnitClassStats["SmallLand"],
		Attack:  0,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"scholar": {
		Name:    "Scholar",
		Class:   UnitClassStats["SmallLand"],
		Attack:  0,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"caravan": {
		Name:    "Caravan",
		Class:   UnitClassStats["Merchant"],
		Attack:  0,
		Defense: 1,
		HP:      10,
		FP:      1,
	},
	"freight": {
		Name:    "Freight",
		Class:   UnitClassStats["Merchant"],
		Attack:  0,
		Defense: 2,
		HP:      20,
		FP:      1,
	},
	"explorer": {
		Name:    "Explorer",
		Class:   UnitClassStats["SmallLand"],
		Attack:  0,
		Defense: 1,
		HP:      7,
		FP:      1,
	},
	"leader": {
		Name:    "Leader",
		Class:   UnitClassStats["Land"],
		Attack:  0,
		Defense: 2,
		HP:      20,
		FP:      1,
	},
	"barbarian_leader": {
		Name:    "Barbarian Leader",
		Class:   UnitClassStats["Land"],
		Attack:  0,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"awacs": {
		Name:    "AWACS",
		Class:   UnitClassStats["Air"],
		Attack:  0,
		Defense: 1,
		HP:      20,
		FP:      1,
	},
	"fusion_fighter": {
		Name:        "Fusion Fighter",
		Class:       UnitClassStats["Air"],
		Attack:      8,
		Defense:     7,
		HP:          20,
		FP:          1,
		AirAttacker: true,
	},
	"fusion_bomber": {
		Name:        "Fusion Bomber",
		Class:       UnitClassStats["Air"],
		Attack:      9,
		Defense:     3,
		HP:          30,
		FP:          1,
		AirAttacker: true,
	},
	"fusion_battleship": {
		Name:            "Fusion Battleship",
		Class:           UnitClassStats["Sea"],
		Attack:          24,
		Defense:         12,
		HP:              30,
		FP:              3,
		BadCityDefender: true,
	},
	"fusion_armor": {
		Name:    "Fusion Armor",
		Class:   UnitClassStats["BigLand"],
		Attack:  10,
		Defense: 5,
		HP:      30,
		FP:      2,
	},
	"cargo_aircraft": {
		Name:    "Cargo Aircraft",
		Class:   UnitClassStats["Air"],
		Attack:  0,
		Defense: 3,
		HP:      30,
		FP:      1,
	},
}
