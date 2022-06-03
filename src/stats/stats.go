package stats

import (
	"github.com/xarxziux/freecoco/src/models.go"
)

func getTerrainType(terrain string) models.TerrainType {
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
}

func getCitySize(size int) models.CityType {
	if size == 0 {
		return models.NoCity
	}

	if size < 9 {
		return models.Town
	}

	return models.City
}

func getCityBonus(city CityStats) float32 {
	citySize := getCitySize(city.Size)

	if citySize == models.NoCity {
		return 1
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

//func getDefenceStats(city CityStats) float32 {
