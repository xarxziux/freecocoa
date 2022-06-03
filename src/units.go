package main

import (
	"github.com/xarxziux/freecoco/src/models.go"
)

/*
 * This file lists all units combat stats as extracted from units.ruleset.
 */
var UnitStats = map[string]UnitDetails{
	"settlers": {
		Name:    "Settlers",
		Class:   models.SmallLand,
		Attack:  0,
		Defense: 1,
		HP:      20,
		FP:      1,
	},
	"migrants": {
		Name:    "Migrants",
		Class:   models.SmallLand,
		Attack:  0,
		Defense: 1,
		HP:      20,
		FP:      1,
	},
	"immigrants": {
		Name:    "Immigrants",
		Class:   models.SmallLand,
		Attack:  0,
		Defense: 1,
		HP:      20,
		FP:      1,
	},
	"tribal_worker": {
		Name:    "Tribal Workers",
		Class:   models.SmallLand,
		Attack:  0,
		Defense: 1,
		HP:      10,
		FP:      1,
	},
	"worker": {
		Name:    "Workers",
		Class:   models.SmallLand,
		Attack:  0,
		Defense: 1,
		HP:      10,
		FP:      1,
	},
	"engineers": {
		Name:    "Engineers",
		Class:   models.SmallLand,
		Attack:  0,
		Defense: 2,
		HP:      20,
		FP:      1,
	},
	"warriors": {
		Name:    "Warriors",
		Class:   models.Land,
		Attack:  1,
		Defense: 1,
		HP:      10,
		FP:      1,
	},
	"phalanx": {
		Name:    "Phalanx",
		Class:   models.Land,
		Attack:  1,
		Defense: 2,
		HP:      10,
		FP:      1,
	},
	"archers": {
		Name:    "Archers",
		Class:   models.Land,
		Attack:  3,
		Defense: 1,
		HP:      10,
		FP:      1,
	},
	"legion": {
		Name:    "Swordsmen",
		Class:   models.Land,
		Attack:  4,
		Defense: 2,
		HP:      10,
		FP:      1,
	},
	"pikemen": {
		Name:    "Pikemen",
		Class:   models.Land,
		Attack:  2,
		Defense: 3,
		HP:      10,
		FP:      1,
	},
	"musketeers": {
		Name:    "Musketeers",
		Class:   models.Land,
		Attack:  3,
		Defense: 3,
		HP:      20,
		FP:      1,
	},
	"riflemen": {
		Name:    "Riflemen",
		Class:   models.Land,
		Attack:  5,
		Defense: 4,
		HP:      20,
		FP:      1,
	},
	"infantry": {
		Name:    "Infantry",
		Class:   models.Land,
		Attack:  5,
		Defense: 6,
		HP:      20,
		FP:      1,
	},
	"alpine_troops": {
		Name:    "Alpine Troops",
		Class:   models.Land,
		Attack:  7,
		Defense: 4,
		HP:      20,
		FP:      1,
	},
	"partisan": {
		Name:    "Partisan",
		Class:   models.Land,
		Attack:  4,
		Defense: 5,
		HP:      20,
		FP:      1,
	},
	"fanatics": {
		Name:    "Fanatics",
		Class:   models.Land,
		Attack:  3,
		Defense: 4,
		HP:      20,
		FP:      1,
	},
	"marines": {
		Name:    "Marines",
		Class:   models.Land,
		Attack:  8,
		Defense: 5,
		HP:      20,
		FP:      1,
	},
	"paratroopers": {
		Name:    "Paratroopers",
		Class:   models.Land,
		Attack:  6,
		Defense: 4,
		HP:      20,
		FP:      1,
	},
	"mech_inf": {
		Class:   models.BigLand,
		Attack:  6,
		Defense: 6,
		HP:      30,
		FP:      1,
	},
	"horsemen": {
		Name:    "Horsemen",
		Class:   models.Land,
		Attack:  2,
		Defense: 1,
		HP:      10,
		FP:      1,
	},
	"chariot": {
		Name:    "Chariot",
		Class:   models.BigLand,
		Attack:  4,
		Defense: 1,
		HP:      10,
		FP:      1,
	},
	"elephants": {
		Name:    "Elephants",
		Class:   models.Land,
		Attack:  3,
		Defense: 2,
		HP:      10,
		FP:      1,
	},
	"crusaders": {
		Name:    "Crusaders",
		Class:   models.Land,
		Attack:  3,
		Defense: 2,
		HP:      15,
		FP:      1,
	},
	"knights": {
		Name:    "Knights",
		Class:   models.Land,
		Attack:  6,
		Defense: 3,
		HP:      10,
		FP:      1,
	},
	"dragoons": {
		Name:    "Dragoons",
		Class:   models.Land,
		Attack:  5,
		Defense: 2,
		HP:      20,
		FP:      1,
	},
	"cavalry": {
		Name:    "Cavalry",
		Class:   models.Land,
		Attack:  8,
		Defense: 3,
		HP:      20,
		FP:      1,
	},
	"armor": {
		Name:    "Armor",
		Class:   models.BigLand,
		Attack:  10,
		Defense: 5,
		HP:      30,
		FP:      1,
	},
	"catapult": {
		Name:       "Catapult",
		Class:      models.BigSiege,
		Attack:     5,
		Defense:    1,
		HP:         10,
		FP:         1,
		CityBuster: true,
	},
	"cannon": {
		Name:       "Cannon",
		Class:      models.BigSiege,
		Attack:     8,
		Defense:    1,
		HP:         20,
		FP:         1,
		CityBuster: true,
	},
	"artillery": {
		Name:       "Artillery",
		Class:      models.BigSiege,
		Attack:     10,
		Defense:    2,
		HP:         20,
		FP:         1,
		CityBuster: true,
	},
	"howitzer": {
		Name:       "Howitzer",
		Class:      models.BigSiege,
		Attack:     12,
		Defense:    2,
		HP:         30,
		FP:         1,
		CityBuster: true,
	},
	"fighter": {
		Name:    "Fighter",
		Class:   models.Air,
		Attack:  4,
		Defense: 4,
		HP:      20,
		FP:      1,
	},
	"bomber": {
		Name:    "Bomber",
		Class:   models.Air,
		Attack:  6,
		Defense: 1,
		HP:      30,
		FP:      1,
	},
	"helicopter": {
		Name:    "Helicopter",
		Class:   models.Helicopter,
		Attack:  5,
		Defense: 3,
		HP:      30,
		FP:      1,
	},
	"stealth_fighter": {
		Name:    "Stealth Fighter",
		Class:   models.Air,
		Attack:  8,
		Defense: 7,
		HP:      20,
		FP:      1,
	},
	"stealth_bomber": {
		Name:    "Stealth Bomber",
		Class:   models.Air,
		Attack:  9,
		Defense: 3,
		HP:      30,
		FP:      1,
	},
	"trireme": {
		Name:    "Trireme",
		Class:   models.Trireme,
		Attack:  1,
		Defense: 1,
		HP:      10,
		FP:      1,
	},
	"barge": {
		Name:    "Barge",
		Class:   models.Trireme,
		Attack:  0,
		Defense: 2,
		HP:      20,
		FP:      1,
	},
	"longboat": {
		Name:    "Longboat",
		Class:   models.Sea,
		Attack:  1,
		Defense: 2,
		HP:      4,
		FP:      1,
	},
	"square_rigged_caravel": {
		Class:   models.Sea,
		Attack:  3,
		Defense: 2,
		HP:      8,
		FP:      1,
	},
	"caravel": {
		Name:    "Caravel",
		Class:   models.Sea,
		Attack:  1,
		Defense: 2,
		HP:      10,
		FP:      1,
	},
	"galleon": {
		Name:    "Galleon",
		Class:   models.Sea,
		Attack:  0,
		Defense: 2,
		HP:      20,
		FP:      1,
	},
	"frigate": {
		Name:    "Frigate",
		Class:   models.Sea,
		Attack:  4,
		Defense: 3,
		HP:      20,
		FP:      1,
	},
	"flagship_frigate": {
		Name:    "Flagship Frigate",
		Class:   models.Sea,
		Attack:  5,
		Defense: 3,
		HP:      22,
		FP:      1,
	},
	"ironclad": {
		Name:    "Ironclad",
		Class:   models.Sea,
		Attack:  5,
		Defense: 6,
		HP:      20,
		FP:      1,
	},
	"destroyer": {
		Name:    "Destroyer",
		Class:   models.Sea,
		Attack:  4,
		Defense: 4,
		HP:      30,
		FP:      1,
	},
	"cruiser": {
		Name:    "Cruiser",
		Class:   models.Sea,
		Attack:  4,
		Defense: 8,
		HP:      30,
		FP:      1,
	},
	"aegis_cruiser": {
		Name:    "AEGIS Cruiser",
		Class:   models.Sea,
		Attack:  8,
		Defense: 8,
		HP:      30,
		FP:      1,
	},
	"battleship": {
		Name:    "Battleship",
		Class:   models.Sea,
		Attack:  14,
		Defense: 14,
		HP:      30,
		FP:      2,
	},
	"submarine": {
		Name:    "Submarine",
		Class:   models.Sea,
		Attack:  12,
		Defense: 4,
		HP:      20,
		FP:      1,
	},
	"nuclear_submarine": {
		Name:    "Nuclear Submarine",
		Class:   models.Sea,
		Attack:  12,
		Defense: 4,
		HP:      20,
		FP:      1,
	},
	"carrier": {
		Name:    "Carrier",
		Class:   models.Sea,
		Attack:  0,
		Defense: 10,
		HP:      40,
		FP:      1,
	},
	"transport": {
		Name:    "Transport",
		Class:   models.Sea,
		Attack:  0,
		Defense: 3,
		HP:      30,
		FP:      1,
	},
	"missile": {
		Name:    "Missile",
		Class:   models.Missile,
		Attack:  15,
		Defense: 0,
		HP:      2,
		FP:      2,
	},
	"cruise_missile": {
		Name:       "Cruise Missile",
		Class:      models.Missile,
		Attack:     10,
		Defense:    0,
		HP:         5,
		FP:         3,
		CityBuster: true,
	},
	"intercontinental_missile": {
		Name:    "Intercontinental Missile",
		Class:   models.Missile,
		Attack:  25,
		Defense: 0,
		HP:      5,
		FP:      3,
	},
	"nuclear_bomb": {
		Name:    "Nuclear Bomb",
		Class:   models.Missile,
		Attack:  99,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"nuclear": {
		Name:    "Nuclear",
		Class:   models.Missile,
		Attack:  99,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"diplomat": {
		Name:    "Diplomat",
		Class:   models.SmallLand,
		Attack:  0,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"spy": {
		Name:    "Spy",
		Class:   models.SmallLand,
		Attack:  0,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"operative": {
		Name:    "Operative",
		Class:   models.SmallLand,
		Attack:  0,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"scribe": {
		Name:    "Scribe",
		Class:   models.SmallLand,
		Attack:  0,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"scholar": {
		Name:    "Scholar",
		Class:   models.SmallLand,
		Attack:  0,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"caravan": {
		Name:    "Caravan",
		Class:   models.Merchant,
		Attack:  0,
		Defense: 1,
		HP:      10,
		FP:      1,
	},
	"freight": {
		Name:    "Freight",
		Class:   models.Merchant,
		Attack:  0,
		Defense: 2,
		HP:      20,
		FP:      1,
	},
	"explorer": {
		Name:    "Explorer",
		Class:   models.SmallLand,
		Attack:  0,
		Defense: 1,
		HP:      7,
		FP:      1,
	},
	"leader": {
		Name:    "Leader",
		Class:   models.Land,
		Attack:  0,
		Defense: 2,
		HP:      20,
		FP:      1,
	},
	"barbarian_leader": {
		Name:    "Barbarian Leader",
		Class:   models.Land,
		Attack:  0,
		Defense: 0,
		HP:      10,
		FP:      1,
	},
	"awacs": {
		Name:    "AWACS",
		Class:   models.Air,
		Attack:  0,
		Defense: 1,
		HP:      20,
		FP:      1,
	},
	"fusion_fighter": {
		Name:    "Fusion Fighter",
		Class:   models.Air,
		Attack:  8,
		Defense: 7,
		HP:      20,
		FP:      1,
	},
	"fusion_bomber": {
		Name:    "Fusion Bomber",
		Class:   models.Air,
		Attack:  9,
		Defense: 3,
		HP:      30,
		FP:      1,
	},
	"fusion_battleship": {
		Name:    "Fusion Battleship",
		Class:   models.Sea,
		Attack:  24,
		Defense: 12,
		HP:      30,
		FP:      3,
	},
	"fusion_armor": {
		Name:    "Fusion Armor",
		Class:   models.BigLand,
		Attack:  10,
		Defense: 5,
		HP:      30,
		FP:      2,
	},
	"cargo_aircraft": {
		Name:    "Cargo Aircraft",
		Class:   models.Air,
		Attack:  0,
		Defense: 3,
		HP:      30,
		FP:      1,
	},
}
