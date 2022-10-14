package core

import (
	"freecocoa/src/models"
)

/*func GetTerrainType(terrain string) models.TerrainType {
	switch terrain {
	case "deep ocean":
	case "desert":
	case "glacier":
	case "grassland":
	case "plains":
	case "tundra":
		return models.Flat
	case "lake":
	case "ocean":
		return models.Shallow
	case "forest":
	case "jungle":
	case "swamp":
		return models.Rough
	case "hills":
		return models.Hills
	case "mountains":
		return models.Mountains
	}

	return models.NoTerrain
}*/

func GetCitySize(size uint8) models.CityType {
	if size == 0 {
		return models.NoCity
	}

	if size < 9 {
		return models.Town
	}

	return models.City
}

func GetCityLandBonus(city models.CityStats) float32 {
	citySize := GetCitySize(city.Size)

	if citySize == models.NoCity {
		return 1.0
	}

	if citySize == models.Town && !city.Walls {
		return 1.5
	}

	if citySize == models.Town && city.Walls {
		return 2.25
	}

	if citySize == models.City && !city.Walls {
		return 2.0
	}

	if citySize == models.City && city.Walls {
		return 3.0
	}

	return 0
}

func GetCitySeaBonus(city models.CityStats) float32 {
	citySize := GetCitySize(city.Size)

	if citySize == models.NoCity {
		return 1
	}

	if citySize == models.Town && !city.CoastalDefence {
		return 1.5
	}

	if citySize == models.Town && city.CoastalDefence {
		return 3.0
	}

	if citySize == models.City && !city.CoastalDefence {
		return 2.0
	}

	if citySize == models.City && city.CoastalDefence {
		return 3.0
	}

	return 0
}

// func getDefenceStats(city CityStats) float32 {
