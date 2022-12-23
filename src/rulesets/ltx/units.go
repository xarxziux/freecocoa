package ltx

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
	"nuclear": {
		NameEnum:    models.Nuclear,
		Missile:     true,
		Unreachable: true,
	},
	"land": {
		NameEnum:       models.Land,
		TerrainDefense: true,
		CanFortify:     true,
	},
	"ancient_land": {
		NameEnum:       models.AncientLand,
		TerrainDefense: true,
		CanFortify:     true,
	},
	"small_unit": {
		NameEnum:       models.SmallUnit,
		TerrainDefense: true,
		CanFortify:     true,
	},
	"smallland": {
		NameEnum:       models.SmallLand,
		TerrainDefense: true,
		CanFortify:     true,
	},
	"amphibious": {
		NameEnum:       models.Amphibious,
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
	"small_sea": {
		NameEnum: models.SmallSea,
	},
	"deep_sea": {
		NameEnum:    models.DeepSea,
		Unreachable: true,
	},
	"trireme": {
		NameEnum: models.Trireme,
	},
	"patrol": {
		NameEnum: models.Patrol,
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
		Name:  "Settlers",
		Class: UnitClassStats["smallland"],
		AP:    0,
		DP:    1,
		HP:    20,
		FP:    1,
	},
	"migrants": {
		Name:  "Migrants",
		Class: UnitClassStats["smallland"],
		AP:    0,
		DP:    1,
		HP:    20,
		FP:    1,
	},
	"immigrants": {
		Name:  "Immigrants",
		Class: UnitClassStats["smallland"],
		AP:    0,
		DP:    1,
		HP:    20,
		FP:    1,
	},
	"tribal_worker": {
		Name:  "Tribal Workers",
		Class: UnitClassStats["smallland"],
		AP:    0,
		DP:    1,
		HP:    12,
		FP:    1,
	},
	"worker": {
		Name:  "Workers",
		Class: UnitClassStats["smallland"],
		AP:    0,
		DP:    1,
		HP:    10,
		FP:    1,
	},
	"engineers": {
		Name:  "Engineers",
		Class: UnitClassStats["smallland"],
		AP:    0,
		DP:    2,
		HP:    20,
		FP:    1,
	},
	"warriors": {
		Name:  "Warriors",
		Class: UnitClassStats["ancientland"],
		AP:    1,
		DP:    1,
		HP:    10,
		FP:    1,
	},
	"phalanx": {
		Name:  "Phalanx",
		Class: UnitClassStats["land"],
		AP:    1,
		DP:    2,
		HP:    10,
		FP:    1,
	},
	"archers": {
		Name:  "Archers",
		Class: UnitClassStats["land"],
		AP:    3,
		DP:    1,
		HP:    10,
		FP:    1,
	},
	"legion": {
		Name:  "Swordsmen",
		Class: UnitClassStats["land"],
		AP:    4,
		DP:    2,
		HP:    10,
		FP:    1,
	},
	"pikemen": {
		Name:  "Pikemen",
		Class: UnitClassStats["land"],
		AP:    2,
		DP:    3,
		HP:    10,
		FP:    1,
	},
	"musketeers": {
		Name:  "Musketeers",
		Class: UnitClassStats["land"],
		AP:    3,
		DP:    3,
		HP:    20,
		FP:    1,
	},
	"navy_troops": {
		Name:  "Navy Troops",
		Class: UnitClassStats["smallunit"],
		AP:    2,
		DP:    2,
		HP:    15,
		FP:    1,
	},
	"riflemen": {
		Name:  "Riflemen",
		Class: UnitClassStats["land"],
		AP:    5,
		DP:    4,
		HP:    20,
		FP:    1,
	},
	"infantry": {
		Name:  "Infantry",
		Class: UnitClassStats["land"],
		AP:    5,
		DP:    6,
		HP:    20,
		FP:    1,
	},
	"alpine_troops": {
		Name:  "Alpine Troops",
		Class: UnitClassStats["land"],
		AP:    7,
		DP:    4,
		HP:    20,
		FP:    1,
	},
	"partisan": {
		Name:  "Partisan",
		Class: UnitClassStats["land"],
		AP:    4,
		DP:    5,
		HP:    20,
		FP:    1,
	},
	"fanatics": {
		Name:  "Fanatics",
		Class: UnitClassStats["land"],
		AP:    3,
		DP:    4,
		HP:    20,
		FP:    1,
	},
	"marines": {
		Name:  "Marines",
		Class: UnitClassStats["land"],
		AP:    8,
		DP:    5,
		HP:    20,
		FP:    1,
	},
	"paratroopers": {
		Name:  "Paratroopers",
		Class: UnitClassStats["land"],
		AP:    6,
		DP:    4,
		HP:    20,
		FP:    1,
	},
	"mech_inf": {
		Name:  "Mech. Inf.",
		Class: UnitClassStats["bigland"],
		AP:    6,
		DP:    6,
		HP:    30,
		FP:    1,
	},
	"horsemen": {
		Name:  "Horsemen",
		Class: UnitClassStats["land"],
		AP:    2,
		DP:    1,
		HP:    10,
		FP:    1,
		Horse: true,
	},
	"chariot": {
		Name:  "Chariot",
		Class: UnitClassStats["bigland"],
		AP:    4,
		DP:    1,
		HP:    10,
		FP:    1,
		Horse: true,
	},
	"elephants": {
		Name:  "Elephants",
		Class: UnitClassStats["land"],
		AP:    3,
		DP:    2,
		HP:    10,
		FP:    1,
		Horse: true,
	},
	"crusaders": {
		Name:  "Crusaders",
		Class: UnitClassStats["land"],
		AP:    3,
		DP:    2,
		HP:    15,
		FP:    1,
		Horse: true,
	},
	"knights": {
		Name:  "Knights",
		Class: UnitClassStats["land"],
		AP:    6,
		DP:    2,
		HP:    10,
		FP:    1,
		Horse: true,
	},
	"dragoons": {
		Name:  "Dragoons",
		Class: UnitClassStats["land"],
		AP:    5,
		DP:    2,
		HP:    20,
		FP:    1,
		Horse: true,
	},
	"cavalry": {
		Name:  "Cavalry",
		Class: UnitClassStats["land"],
		AP:    8,
		DP:    3,
		HP:    20,
		FP:    1,
		Horse: true,
	},
	"armor": {
		Name:  "Armor",
		Class: UnitClassStats["bigland"],
		AP:    10,
		DP:    5,
		HP:    30,
		FP:    1,
	},
	"catapult": {
		Name:       "Catapult",
		Class:      UnitClassStats["bigsiege"],
		AP:         5,
		DP:         1,
		HP:         10,
		FP:         1,
		CityBuster: true,
	},
	"trebuchet": {
		Name:       "Trebuchet",
		Class:      UnitClassStats["bigsiege"],
		AP:         10,
		DP:         1,
		HP:         10,
		FP:         1,
		CityBuster: true,
	},
	"cannon": {
		Name:       "Cannon",
		Class:      UnitClassStats["bigsiege"],
		AP:         8,
		DP:         1,
		HP:         20,
		FP:         1,
		CityBuster: true,
	},
	"artillery": {
		Name:       "Artillery",
		Class:      UnitClassStats["bigsiege"],
		AP:         10,
		DP:         2,
		HP:         20,
		FP:         1,
		CityBuster: true,
	},
	"howitzer": {
		Name:       "Howitzer",
		Class:      UnitClassStats["bigsiege"],
		AP:         12,
		DP:         2,
		HP:         30,
		FP:         1,
		CityBuster: true,
	},
	"fighter": {
		Name:        "Fighter",
		Class:       UnitClassStats["air"],
		AP:          4,
		DP:          4,
		HP:          20,
		FP:          1,
		AirAttacker: true,
	},
	"bomber": {
		Name:        "Bomber",
		Class:       UnitClassStats["air"],
		AP:          6,
		DP:          1,
		HP:          30,
		FP:          1,
		AirAttacker: true,
	},
	"helicopter": {
		Name:        "Helicopter",
		Class:       UnitClassStats["helicopter"],
		AP:          5,
		DP:          3,
		HP:          30,
		FP:          1,
		AirAttacker: true,
	},
	"stealth_fighter": {
		Name:        "Stealth Fighter",
		Class:       UnitClassStats["air"],
		AP:          8,
		DP:          7,
		HP:          20,
		FP:          1,
		AirAttacker: true,
	},
	"stealth_bomber": {
		Name:        "Stealth Bomber",
		Class:       UnitClassStats["air"],
		AP:          9,
		DP:          3,
		HP:          30,
		FP:          1,
		AirAttacker: true,
	},
	"trireme": {
		Name:            "Trireme",
		Class:           UnitClassStats["trireme"],
		AP:              1,
		DP:              1,
		HP:              10,
		FP:              1,
		BadCityDefender: true,
	},
	"barge": {
		Name:            "Barge",
		Class:           UnitClassStats["trireme"],
		AP:              0,
		DP:              2,
		HP:              20,
		FP:              1,
		BadCityDefender: true,
	},
	"longboat": {
		Name:            "Longboat",
		Class:           UnitClassStats["sea"],
		AP:              1,
		DP:              2,
		HP:              4,
		FP:              1,
		BadCityDefender: true,
	},
	"war_longboat": {
		Name:            "War Longboat",
		Class:           UnitClassStats["sea"],
		AP:              1,
		DP:              1,
		HP:              1,
		FP:              1,
		BadCityDefender: true,
	},
	"pirogue": {
		Name:            "Pirogue",
		Class:           UnitClassStats["sea"],
		AP:              0,
		DP:              1,
		HP:              5,
		FP:              1,
		BadCityDefender: true,
	},
	"cog": {
		Name:            "Cog",
		Class:           UnitClassStats["patrol"],
		AP:              0,
		DP:              2,
		HP:              8,
		FP:              1,
		BadCityDefender: true,
	},
	"patrol_cutter": {
		Name:            "Patrol Cutter",
		Class:           UnitClassStats["patrol"],
		AP:              0,
		DP:              2,
		HP:              20,
		FP:              1,
		BadCityDefender: true,
	},
	"square_rigged_caravel": {
		Class:           UnitClassStats["sea"],
		AP:              2,
		DP:              2,
		HP:              12,
		FP:              1,
		BadCityDefender: true,
	},
	"caravel": {
		Name:            "Caravel",
		Class:           UnitClassStats["sea"],
		AP:              2,
		DP:              2,
		HP:              10,
		FP:              1,
		BadCityDefender: true,
	},
	"galleon": {
		Name:            "Galleon",
		Class:           UnitClassStats["sea"],
		AP:              0,
		DP:              2,
		HP:              20,
		FP:              1,
		BadCityDefender: true,
	},
	"frigate": {
		Name:            "Frigate",
		Class:           UnitClassStats["sea"],
		AP:              4,
		DP:              3,
		HP:              20,
		FP:              1,
		BadCityDefender: true,
	},
	"flagship_frigate": {
		Name:            "Flagship Frigate",
		Class:           UnitClassStats["sea"],
		AP:              5,
		DP:              3,
		HP:              22,
		FP:              1,
		BadCityDefender: true,
	},
	"ironclad": {
		Name:            "Ironclad",
		Class:           UnitClassStats["sea"],
		AP:              5,
		DP:              6,
		HP:              20,
		FP:              1,
		BadCityDefender: true,
	},
	"destroyer": {
		Name:            "Destroyer",
		Class:           UnitClassStats["sea"],
		AP:              4,
		DP:              4,
		HP:              30,
		FP:              1,
		BadCityDefender: true,
	},
	"cruiser": {
		Name:            "Cruiser",
		Class:           UnitClassStats["sea"],
		AP:              4,
		DP:              8,
		HP:              30,
		FP:              1,
		BadCityDefender: true,
	},
	"aegis_cruiser": {
		Name:            "AEGIS Cruiser",
		Class:           UnitClassStats["sea"],
		AP:              8,
		DP:              8,
		HP:              30,
		FP:              1,
		BadCityDefender: true,
	},
	"battleship": {
		Name:            "Battleship",
		Class:           UnitClassStats["sea"],
		AP:              14,
		DP:              14,
		HP:              30,
		FP:              2,
		BadCityDefender: true,
	},
	"submarine": {
		Name:            "Submarine",
		Class:           UnitClassStats["sea"],
		AP:              12,
		DP:              4,
		HP:              20,
		FP:              1,
		Submarine:       true,
		BadCityDefender: true,
	},
	"nuclear_submarine": {
		Name:            "Nuclear Submarine",
		Class:           UnitClassStats["sea"],
		AP:              12,
		DP:              4,
		HP:              20,
		FP:              1,
		Submarine:       true,
		BadCityDefender: true,
	},
	"nuclear_submarine_type_ii": {
		Name:            "Nuclear Submarine Type II",
		Class:           UnitClassStats["deepsea"],
		AP:              10,
		DP:              6,
		HP:              20,
		FP:              1,
		Submarine:       true,
		BadCityDefender: true,
	},
	"carrier": {
		Name:            "Carrier",
		Class:           UnitClassStats["sea"],
		AP:              0,
		DP:              10,
		HP:              40,
		FP:              1,
		BadCityDefender: true,
	},
	"transport": {
		Name:            "Transport",
		Class:           UnitClassStats["sea"],
		AP:              0,
		DP:              3,
		HP:              30,
		FP:              1,
		BadCityDefender: true,
	},
	"advanced_torpedo": {
		Name:  "Advanced Torpedo",
		Class: UnitClassStats["smallsea"],
		AP:    12,
		DP:    0,
		HP:    2,
		FP:    2,
	},
	"missile": {
		Name:        "Air Defense Missile",
		Class:       UnitClassStats["missile"],
		AP:          15,
		DP:          0,
		HP:          2,
		FP:          2,
		AirAttacker: true,
	},
	"cruise_missile": {
		Name:        "Cruise Missile",
		Class:       UnitClassStats["missile"],
		AP:          10,
		DP:          0,
		HP:          5,
		FP:          3,
		CityBuster:  true,
		AirAttacker: true,
	},
	"intercontinental_missile": {
		Name:        "Intercontinental Missile",
		Class:       UnitClassStats["missile"],
		AP:          25,
		DP:          0,
		HP:          5,
		FP:          3,
		AirAttacker: true,
	},
	"nuclear_bomb": {
		Name:  "Nuclear Bomb",
		Class: UnitClassStats["nuclear"],
		AP:    99,
		DP:    0,
		HP:    10,
		FP:    1,
	},
	"nuclear": {
		Name:        "Nuclear",
		Class:       UnitClassStats["nuclear"],
		AP:          99,
		DP:          0,
		HP:          10,
		FP:          1,
		AirAttacker: true,
	},
	"diplomat": {
		Name:  "Diplomat",
		Class: UnitClassStats["smallland"],
		AP:    0,
		DP:    0,
		HP:    10,
		FP:    1,
	},
	"spy": {
		Name:  "Spy",
		Class: UnitClassStats["smallland"],
		AP:    0,
		DP:    0,
		HP:    10,
		FP:    1,
	},
	"operative": {
		Name:  "Operative",
		Class: UnitClassStats["smallland"],
		AP:    0,
		DP:    0,
		HP:    10,
		FP:    1,
	},
	"scribe": {
		Name:  "Scribe",
		Class: UnitClassStats["smallland"],
		AP:    0,
		DP:    0,
		HP:    10,
		FP:    1,
	},
	"scholar": {
		Name:  "Scholar",
		Class: UnitClassStats["smallland"],
		AP:    0,
		DP:    0,
		HP:    10,
		FP:    1,
	},
	"inventor": {
		Name:  "Inventor",
		Class: UnitClassStats["smallland"],
		AP:    0,
		DP:    0,
		HP:    10,
		FP:    1,
	},
	"caravan": {
		Name:  "Caravan",
		Class: UnitClassStats["merchant"],
		AP:    0,
		DP:    1,
		HP:    10,
		FP:    1,
	},
	"freight": {
		Name:  "Freight",
		Class: UnitClassStats["merchant"],
		AP:    0,
		DP:    2,
		HP:    20,
		FP:    1,
	},
	"explorer": {
		Name:  "Explorer",
		Class: UnitClassStats["smallland"],
		AP:    0,
		DP:    1,
		HP:    7,
		FP:    1,
	},
	"leader": {
		Name:  "Leader",
		Class: UnitClassStats["land"],
		AP:    0,
		DP:    2,
		HP:    20,
		FP:    1,
	},
	"barbarian_leader": {
		Name:  "Barbarian Leader",
		Class: UnitClassStats["land"],
		AP:    0,
		DP:    0,
		HP:    10,
		FP:    1,
	},
	"awacs": {
		Name:  "AWACS",
		Class: UnitClassStats["air"],
		AP:    0,
		DP:    1,
		HP:    20,
		FP:    1,
	},
	"fusion_fighter": {
		Name:        "Fusion Fighter",
		Class:       UnitClassStats["air"],
		AP:          8,
		DP:          7,
		HP:          20,
		FP:          2,
		AirAttacker: true,
	},
	"fusion_bomber": {
		Name:        "Fusion Bomber",
		Class:       UnitClassStats["air"],
		AP:          9,
		DP:          3,
		HP:          30,
		FP:          1,
		AirAttacker: true,
	},
	"fusion_battleship": {
		Name:            "Fusion Battleship",
		Class:           UnitClassStats["sea"],
		AP:              24,
		DP:              12,
		HP:              30,
		FP:              3,
		BadCityDefender: true,
	},
	"fusion_armor": {
		Name:  "Fusion Armor",
		Class: UnitClassStats["bigland"],
		AP:    10,
		DP:    5,
		HP:    30,
		FP:    2,
	},
	"cargo_aircraft": {
		Name:  "Cargo Aircraft",
		Class: UnitClassStats["air"],
		AP:    0,
		DP:    3,
		HP:    30,
		FP:    1,
	},
}
