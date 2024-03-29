package ltt

// AUTO-GENERATED FILE - DO NOT EDIT
// USE THE SCRIPT parse_units_ruleset.js TO RE-GENERATE IF NEEDED

import (
	"freecocoa/src/models"
)

// UnitClassStats lists all unit classes and combat-relevant flags as extracted
// from units.ruleset.
var UnitClassStats = map[string]models.UnitClass{
	"missile": {
		NameEnum:    models.Missile,
		Missile:     true,
		Unreachable: true,
	},
	"land": {
		NameEnum:       models.Land,
		TerrainDefense: true,
		CanFortify:     true,
	},
	"smallland": {
		NameEnum:       models.SmallLand,
		TerrainDefense: true,
		CanFortify:     true,
	},
	"bigland": {
		NameEnum:   models.BigLand,
		CanFortify: true,
	},
	"bigsiege": {
		NameEnum:        models.BigSiege,
		CanFortify:      true,
		AttackNonNative: true,
	},
	"merchant": {
		NameEnum:       models.Merchant,
		TerrainDefense: true,
	},
	"sea": {
		NameEnum:        models.Sea,
		AttackNonNative: true,
	},
	"trireme": {
		NameEnum: models.Trireme,
	},
	"helicopter": {
		NameEnum:    models.Helicopter,
		Unreachable: true,
	},
	"air": {
		NameEnum:    models.Air,
		Unreachable: true,
	},
}

// UnitStats lists all units combat stats and combat-relevant flags as extracted
// from units.ruleset.
var UnitStats = map[string]models.UnitDetails{
	"settlers": {
		Name:      "Settlers",
		Class:     UnitClassStats["smallland"],
		BuildCost: 30,
		AP:        0,
		DP:        1,
		HP:        20,
		FP:        1,
		MP:        27,
	},
	"migrants": {
		Name:      "Migrants",
		Class:     UnitClassStats["smallland"],
		BuildCost: 15,
		AP:        0,
		DP:        1,
		HP:        20,
		FP:        1,
		MP:        27,
	},
	"immigrants": {
		Name:      "Immigrants",
		Class:     UnitClassStats["smallland"],
		BuildCost: 20,
		AP:        0,
		DP:        1,
		HP:        20,
		FP:        1,
		MP:        27,
	},
	"tribal_worker": {
		Name:      "Tribal Workers",
		Class:     UnitClassStats["smallland"],
		BuildCost: 25,
		AP:        0,
		DP:        1,
		HP:        10,
		FP:        1,
		MP:        18,
	},
	"worker": {
		Name:      "Workers",
		Class:     UnitClassStats["smallland"],
		BuildCost: 20,
		AP:        0,
		DP:        1,
		HP:        10,
		FP:        1,
		MP:        27,
	},
	"engineers": {
		Name:      "Engineers",
		Class:     UnitClassStats["smallland"],
		BuildCost: 40,
		AP:        0,
		DP:        2,
		HP:        20,
		FP:        1,
		MP:        54,
	},
	"warriors": {
		Name:      "Warriors",
		Class:     UnitClassStats["land"],
		BuildCost: 10,
		AP:        1,
		DP:        1,
		HP:        10,
		FP:        1,
		MP:        18,
	},
	"phalanx": {
		Name:      "Phalanx",
		Class:     UnitClassStats["land"],
		BuildCost: 15,
		AP:        1,
		DP:        2,
		HP:        10,
		FP:        1,
		MP:        18,
	},
	"archers": {
		Name:      "Archers",
		Class:     UnitClassStats["land"],
		BuildCost: 20,
		AP:        3,
		DP:        1,
		HP:        10,
		FP:        1,
		MP:        18,
	},
	"legion": {
		Name:      "Swordsmen",
		Class:     UnitClassStats["land"],
		BuildCost: 30,
		AP:        4,
		DP:        2,
		HP:        10,
		FP:        1,
		MP:        18,
	},
	"pikemen": {
		Name:      "Pikemen",
		Class:     UnitClassStats["land"],
		BuildCost: 25,
		AP:        2,
		DP:        3,
		HP:        10,
		FP:        1,
		MP:        27,
	},
	"musketeers": {
		Name:      "Musketeers",
		Class:     UnitClassStats["land"],
		BuildCost: 40,
		AP:        3,
		DP:        3,
		HP:        20,
		FP:        1,
		MP:        27,
	},
	"riflemen": {
		Name:      "Riflemen",
		Class:     UnitClassStats["land"],
		BuildCost: 50,
		AP:        5,
		DP:        4,
		HP:        20,
		FP:        1,
		MP:        27,
	},
	"infantry": {
		Name:      "Infantry",
		Class:     UnitClassStats["land"],
		BuildCost: 50,
		AP:        5,
		DP:        6,
		HP:        20,
		FP:        1,
		MP:        27,
	},
	"alpine_troops": {
		Name:      "Alpine Troops",
		Class:     UnitClassStats["land"],
		BuildCost: 60,
		AP:        7,
		DP:        4,
		HP:        20,
		FP:        1,
		MP:        18,
	},
	"partisan": {
		Name:      "Partisan",
		Class:     UnitClassStats["land"],
		BuildCost: 80,
		AP:        4,
		DP:        5,
		HP:        20,
		FP:        1,
		MP:        18,
	},
	"fanatics": {
		Name:      "Fanatics",
		Class:     UnitClassStats["land"],
		BuildCost: 35,
		AP:        3,
		DP:        4,
		HP:        20,
		FP:        1,
		MP:        27,
	},
	"marines": {
		Name:      "Marines",
		Class:     UnitClassStats["land"],
		BuildCost: 60,
		AP:        8,
		DP:        5,
		HP:        20,
		FP:        1,
		MP:        27,
	},
	"paratroopers": {
		Name:      "Paratroopers",
		Class:     UnitClassStats["land"],
		BuildCost: 60,
		AP:        6,
		DP:        4,
		HP:        20,
		FP:        1,
		MP:        27,
	},
	"mech_inf": {
		Name:      "Mech. Inf.",
		Class:     UnitClassStats["bigland"],
		BuildCost: 60,
		AP:        6,
		DP:        6,
		HP:        30,
		FP:        1,
		MP:        81,
	},
	"horsemen": {
		Name:      "Horsemen",
		Class:     UnitClassStats["land"],
		BuildCost: 18,
		AP:        2,
		DP:        1,
		HP:        10,
		FP:        1,
		MP:        36,
		Horse:     true,
	},
	"chariot": {
		Name:      "Chariot",
		Class:     UnitClassStats["bigland"],
		BuildCost: 28,
		AP:        4,
		DP:        1,
		HP:        10,
		FP:        1,
		MP:        36,
		Horse:     true,
	},
	"elephants": {
		Name:      "Elephants",
		Class:     UnitClassStats["land"],
		BuildCost: 28,
		AP:        3,
		DP:        2,
		HP:        10,
		FP:        1,
		MP:        36,
		Horse:     true,
	},
	"crusaders": {
		Name:      "Crusaders",
		Class:     UnitClassStats["land"],
		BuildCost: 28,
		AP:        3,
		DP:        2,
		HP:        15,
		FP:        1,
		MP:        45,
		Horse:     true,
	},
	"knights": {
		Name:      "Knights",
		Class:     UnitClassStats["land"],
		BuildCost: 32,
		AP:        6,
		DP:        3,
		HP:        10,
		FP:        1,
		MP:        36,
		Horse:     true,
	},
	"dragoons": {
		Name:      "Dragoons",
		Class:     UnitClassStats["land"],
		BuildCost: 50,
		AP:        5,
		DP:        2,
		HP:        20,
		FP:        1,
		MP:        54,
		Horse:     true,
	},
	"cavalry": {
		Name:      "Cavalry",
		Class:     UnitClassStats["land"],
		BuildCost: 60,
		AP:        8,
		DP:        3,
		HP:        20,
		FP:        1,
		MP:        54,
		Horse:     true,
	},
	"armor": {
		Name:      "Armor",
		Class:     UnitClassStats["bigland"],
		BuildCost: 80,
		AP:        10,
		DP:        5,
		HP:        30,
		FP:        1,
		MP:        81,
	},
	"catapult": {
		Name:       "Catapult",
		Class:      UnitClassStats["bigsiege"],
		BuildCost:  30,
		AP:         5,
		DP:         1,
		HP:         10,
		FP:         1,
		MP:         27,
		CityBuster: true,
	},
	"cannon": {
		Name:       "Cannon",
		Class:      UnitClassStats["bigsiege"],
		BuildCost:  50,
		AP:         8,
		DP:         1,
		HP:         20,
		FP:         1,
		MP:         27,
		CityBuster: true,
	},
	"artillery": {
		Name:       "Artillery",
		Class:      UnitClassStats["bigsiege"],
		BuildCost:  60,
		AP:         10,
		DP:         2,
		HP:         20,
		FP:         1,
		MP:         36,
		CityBuster: true,
	},
	"howitzer": {
		Name:       "Howitzer",
		Class:      UnitClassStats["bigsiege"],
		BuildCost:  70,
		AP:         12,
		DP:         2,
		HP:         30,
		FP:         1,
		MP:         81,
		CityBuster: true,
	},
	"fighter": {
		Name:        "Fighter",
		Class:       UnitClassStats["air"],
		BuildCost:   40,
		AP:          4,
		DP:          4,
		HP:          20,
		FP:          1,
		MP:          108,
		AirAttacker: true,
	},
	"bomber": {
		Name:        "Bomber",
		Class:       UnitClassStats["air"],
		BuildCost:   90,
		AP:          6,
		DP:          1,
		HP:          30,
		FP:          1,
		MP:          216,
		AirAttacker: true,
	},
	"helicopter": {
		Name:        "Helicopter",
		Class:       UnitClassStats["helicopter"],
		BuildCost:   50,
		AP:          5,
		DP:          3,
		HP:          30,
		FP:          1,
		MP:          108,
		AirAttacker: true,
	},
	"stealth_fighter": {
		Name:        "Stealth Fighter",
		Class:       UnitClassStats["air"],
		BuildCost:   80,
		AP:          8,
		DP:          7,
		HP:          20,
		FP:          1,
		MP:          162,
		AirAttacker: true,
	},
	"stealth_bomber": {
		Name:        "Stealth Bomber",
		Class:       UnitClassStats["air"],
		BuildCost:   120,
		AP:          9,
		DP:          3,
		HP:          30,
		FP:          1,
		MP:          324,
		AirAttacker: true,
	},
	"trireme": {
		Name:            "Trireme",
		Class:           UnitClassStats["trireme"],
		BuildCost:       30,
		AP:              1,
		DP:              1,
		HP:              10,
		FP:              1,
		MP:              72,
		BadCityDefender: true,
	},
	"barge": {
		Name:            "Barge",
		Class:           UnitClassStats["trireme"],
		BuildCost:       45,
		AP:              0,
		DP:              2,
		HP:              20,
		FP:              1,
		MP:              81,
		BadCityDefender: true,
	},
	"longboat": {
		Name:            "Longboat",
		Class:           UnitClassStats["sea"],
		BuildCost:       30,
		AP:              1,
		DP:              2,
		HP:              4,
		FP:              1,
		MP:              72,
		BadCityDefender: true,
	},
	"square_rigged_caravel": {
		Class:           UnitClassStats["sea"],
		BuildCost:       55,
		AP:              3,
		DP:              2,
		HP:              8,
		FP:              1,
		MP:              81,
		BadCityDefender: true,
	},
	"caravel": {
		Name:            "Caravel",
		Class:           UnitClassStats["sea"],
		BuildCost:       40,
		AP:              1,
		DP:              2,
		HP:              10,
		FP:              1,
		MP:              81,
		BadCityDefender: true,
	},
	"galleon": {
		Name:            "Galleon",
		Class:           UnitClassStats["sea"],
		BuildCost:       50,
		AP:              0,
		DP:              2,
		HP:              20,
		FP:              1,
		MP:              108,
		BadCityDefender: true,
	},
	"frigate": {
		Name:            "Frigate",
		Class:           UnitClassStats["sea"],
		BuildCost:       50,
		AP:              4,
		DP:              3,
		HP:              20,
		FP:              1,
		MP:              117,
		BadCityDefender: true,
	},
	"flagship_frigate": {
		Name:            "Flagship Frigate",
		Class:           UnitClassStats["sea"],
		BuildCost:       70,
		AP:              5,
		DP:              3,
		HP:              22,
		FP:              1,
		MP:              135,
		BadCityDefender: true,
	},
	"ironclad": {
		Name:            "Ironclad",
		Class:           UnitClassStats["sea"],
		BuildCost:       50,
		AP:              5,
		DP:              6,
		HP:              20,
		FP:              1,
		MP:              90,
		BadCityDefender: true,
	},
	"destroyer": {
		Name:            "Destroyer",
		Class:           UnitClassStats["sea"],
		BuildCost:       70,
		AP:              4,
		DP:              4,
		HP:              30,
		FP:              1,
		MP:              162,
		BadCityDefender: true,
	},
	"cruiser": {
		Name:            "Cruiser",
		Class:           UnitClassStats["sea"],
		BuildCost:       75,
		AP:              4,
		DP:              8,
		HP:              30,
		FP:              1,
		MP:              135,
		BadCityDefender: true,
	},
	"aegis_cruiser": {
		Name:            "AEGIS Cruiser",
		Class:           UnitClassStats["sea"],
		BuildCost:       90,
		AP:              8,
		DP:              8,
		HP:              30,
		FP:              1,
		MP:              135,
		BadCityDefender: true,
	},
	"battleship": {
		Name:            "Battleship",
		Class:           UnitClassStats["sea"],
		BuildCost:       180,
		AP:              14,
		DP:              14,
		HP:              30,
		FP:              2,
		MP:              108,
		BadCityDefender: true,
	},
	"submarine": {
		Name:            "Submarine",
		Class:           UnitClassStats["sea"],
		BuildCost:       50,
		AP:              12,
		DP:              4,
		HP:              20,
		FP:              1,
		MP:              108,
		Submarine:       true,
		BadCityDefender: true,
	},
	"nuclear_submarine": {
		Name:            "Nuclear Submarine",
		Class:           UnitClassStats["sea"],
		BuildCost:       120,
		AP:              12,
		DP:              4,
		HP:              20,
		FP:              1,
		MP:              126,
		Submarine:       true,
		BadCityDefender: true,
	},
	"carrier": {
		Name:            "Carrier",
		Class:           UnitClassStats["sea"],
		BuildCost:       120,
		AP:              0,
		DP:              10,
		HP:              40,
		FP:              1,
		MP:              135,
		BadCityDefender: true,
	},
	"transport": {
		Name:            "Transport",
		Class:           UnitClassStats["sea"],
		BuildCost:       60,
		AP:              0,
		DP:              3,
		HP:              30,
		FP:              1,
		MP:              135,
		BadCityDefender: true,
	},
	"missile": {
		Name:        "Missile",
		Class:       UnitClassStats["missile"],
		BuildCost:   45,
		AP:          15,
		DP:          0,
		HP:          2,
		FP:          2,
		MP:          81,
		AirAttacker: true,
	},
	"cruise_missile": {
		Name:        "Cruise Missile",
		Class:       UnitClassStats["missile"],
		BuildCost:   45,
		AP:          10,
		DP:          0,
		HP:          5,
		FP:          3,
		MP:          162,
		CityBuster:  true,
		AirAttacker: true,
	},
	"intercontinental_missile": {
		Name:        "Intercontinental Missile",
		Class:       UnitClassStats["missile"],
		BuildCost:   90,
		AP:          25,
		DP:          0,
		HP:          5,
		FP:          3,
		MP:          315,
		AirAttacker: true,
	},
	"nuclear_bomb": {
		Name:      "Nuclear Bomb",
		Class:     UnitClassStats["missile"],
		BuildCost: 200,
		AP:        99,
		DP:        0,
		HP:        10,
		FP:        1,
		MP:        0,
	},
	"nuclear": {
		Name:        "Nuclear",
		Class:       UnitClassStats["missile"],
		BuildCost:   160,
		AP:          99,
		DP:          0,
		HP:          10,
		FP:          1,
		MP:          0,
		AirAttacker: true,
	},
	"diplomat": {
		Name:      "Diplomat",
		Class:     UnitClassStats["smallland"],
		BuildCost: 30,
		AP:        0,
		DP:        0,
		HP:        10,
		FP:        1,
		MP:        45,
	},
	"spy": {
		Name:      "Spy",
		Class:     UnitClassStats["smallland"],
		BuildCost: 50,
		AP:        0,
		DP:        0,
		HP:        10,
		FP:        1,
		MP:        54,
	},
	"operative": {
		Name:      "Operative",
		Class:     UnitClassStats["smallland"],
		BuildCost: 120,
		AP:        0,
		DP:        0,
		HP:        10,
		FP:        1,
		MP:        27,
	},
	"scribe": {
		Name:      "Scribe",
		Class:     UnitClassStats["smallland"],
		BuildCost: 40,
		AP:        0,
		DP:        0,
		HP:        10,
		FP:        1,
		MP:        36,
	},
	"scholar": {
		Name:      "Scholar",
		Class:     UnitClassStats["smallland"],
		BuildCost: 100,
		AP:        0,
		DP:        0,
		HP:        10,
		FP:        1,
		MP:        36,
	},
	"caravan": {
		Name:      "Caravan",
		Class:     UnitClassStats["merchant"],
		BuildCost: 50,
		AP:        0,
		DP:        1,
		HP:        10,
		FP:        1,
		MP:        27,
	},
	"freight": {
		Name:      "Freight",
		Class:     UnitClassStats["merchant"],
		BuildCost: 50,
		AP:        0,
		DP:        2,
		HP:        20,
		FP:        1,
		MP:        54,
	},
	"explorer": {
		Name:      "Explorer",
		Class:     UnitClassStats["smallland"],
		BuildCost: 30,
		AP:        0,
		DP:        1,
		HP:        7,
		FP:        1,
		MP:        27,
	},
	"leader": {
		Name:      "Leader",
		Class:     UnitClassStats["land"],
		BuildCost: 10,
		AP:        0,
		DP:        2,
		HP:        20,
		FP:        1,
		MP:        18,
	},
	"barbarian_leader": {
		Name:      "Barbarian Leader",
		Class:     UnitClassStats["land"],
		BuildCost: 40,
		AP:        0,
		DP:        0,
		HP:        10,
		FP:        1,
		MP:        18,
	},
	"awacs": {
		Name:      "AWACS",
		Class:     UnitClassStats["air"],
		BuildCost: 140,
		AP:        0,
		DP:        1,
		HP:        20,
		FP:        1,
		MP:        315,
	},
	"fusion_fighter": {
		Name:        "Fusion Fighter",
		Class:       UnitClassStats["air"],
		BuildCost:   200,
		AP:          8,
		DP:          7,
		HP:          20,
		FP:          1,
		MP:          252,
		AirAttacker: true,
	},
	"fusion_bomber": {
		Name:        "Fusion Bomber",
		Class:       UnitClassStats["air"],
		BuildCost:   240,
		AP:          9,
		DP:          3,
		HP:          30,
		FP:          1,
		MP:          315,
		AirAttacker: true,
	},
	"fusion_battleship": {
		Name:            "Fusion Battleship",
		Class:           UnitClassStats["sea"],
		BuildCost:       300,
		AP:              24,
		DP:              12,
		HP:              30,
		FP:              3,
		MP:              144,
		BadCityDefender: true,
	},
	"fusion_armor": {
		Name:      "Fusion Armor",
		Class:     UnitClassStats["bigland"],
		BuildCost: 120,
		AP:        10,
		DP:        5,
		HP:        30,
		FP:        2,
		MP:        81,
	},
	"cargo_aircraft": {
		Name:      "Cargo Aircraft",
		Class:     UnitClassStats["air"],
		BuildCost: 200,
		AP:        0,
		DP:        3,
		HP:        30,
		FP:        1,
		MP:        252,
	},
}
