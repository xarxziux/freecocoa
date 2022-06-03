package main

func getTerrainType(terrain string) TerrainType {
	switch terrain {
	case "deep ocean":
	case "desert":
	case "glacier":
	case "grassland":
	case "plains":
	case "tundra":
		return Flat
	case "lake":
	case "ocean":
		return Shallow
	case "forest":
	case "jungle":
	case "swamp":
		return Rough
	case "hills":
		return Hills
	case "mountains":
		return Mountains
	}

	return NoTerrain
}

func getCitySize(size int) CityType {
	if size == 0 {
		return NoCity
	}

	if size < 9 {
		return Town
	}

	return City
}

func getCityBonus(city CityStats) float32 {
	citySize := getCitySize(city.Size)

	if citySize == NoCity {
		return 1
	}

	if citySize == Town && !city.Walls {
		return 1.5
	}

	if citySize == Town && city.Walls {
		return 2.25
	}

	if citySize == City && !city.Walls {
		return 2.0
	}

	if citySize == City && city.Walls {
		return 3.0
	}

	return 0
}

//func getDefenceStats(city CityStats) float32 {
