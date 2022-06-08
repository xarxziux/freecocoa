package models

type CityType int
type UnitClassName int
type TerrainClass int

const (
	NoCity CityType = iota
	Town
	City
)

const (
	NoUnit UnitClassName = iota
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

const (
	NoTerrain TerrainClass = iota
	LandClass
	OceanicClass
)

var VeteranLevels = [10]float32{1, 1.5, 1.75, 2, 2.25, 2.5, 2.75, 3, 3.25, 3.5}

type TerrainType struct {
	Name         string
	Class        TerrainClass
	DefenseBonus int
	CanHaveRiver bool
	NoCities     bool
	NoFortify    bool
	UnsafeCoast  bool
}

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

type UnitClass struct {
	Name             UnitClassName
	TerrainDefense   bool
	Missile          bool
	Unreachable      bool
	CanFortify       bool
	AttackNonNative  bool
	AttFromNonNative bool
}

type UnitDetails struct {
	Name             string
	Class            UnitClass
	Attack           int
	Defense          int
	HP               int
	FP               int
	CityBuster       bool
	AirAttacker      bool
	Horse            bool
	Submarine        bool
	BadCityDefender  bool
	OnlyNativeAttack bool
}
