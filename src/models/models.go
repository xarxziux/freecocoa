package models

type TerrainType int
type CityType int
type UnitType int

const (
	NoTerrain TerrainType = iota
	Flat
	Shallow
	Rough
	Hills
	Mountains
)

const (
	NoCity CityType = iota
	Town
	City
)

const (
	NoUnit UnitType = iota
	Air
	BigLand
	BigSiege
	Helicopter
	Land
	Merchant
	Missile
	Sea
	SmallLand
	Trireme
)

var VeteranLevels = [10]float32{1, 1.5, 1.75, 2, 2.25, 2.5, 2.75, 3, 3.25, 3.5}

type CityStats struct {
	Size       int  `json:"size"`
	Walls      bool `json:"walls"`
	GreatWalls bool `json:"greatWalls"`
}

type DefenceStats struct {
	Unit      string    `json:"unit"`
	VetLevel  int       `json:"vetLevel"`
	Terrain   string    `json:"terrain"`
	City      CityStats `json:"city"`
	River     bool      `json:"river"`
	Fortified bool      `json:"fortified"`
	Fortress  bool      `json:"fortress"`
}

type UnitDetails struct {
	Name       string
	Class      UnitType
	Attack     int
	Defense    int
	HP         int
	FP         int
	CityBuster bool
}
