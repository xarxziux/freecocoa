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
	NoClass UnitClassName = iota
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

func (c UnitClassName) ToString() string {
	return []string{
		"NoClass",
		"Air",
		"BigLand",
		"BigSiege",
		"Helicopter",
		"Land",
		"Merchant",
		"Missile",
		"Sea",
		"SmallLand",
		"Trireme",
	}[c]
}

const (
	NoTerrain TerrainClass = iota
	LandClass
	OceanicClass
)

var VeteranLevels = [10]float32{1.0, 1.5, 1.75, 2, 2.25, 2.5, 2.75, 3, 3.25, 3.5}

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
	Size           int  `json:"size"`
	Walls          bool `json:"walls"`
	GreatWalls     bool `json:"greatWalls"`
	CoastalDefence bool `json:"coastalDefence"`
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
	NameEnum         UnitClassName `json:"-"`
	Name             string
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
	Attack           float32
	Defense          float32
	HP               int
	FP               int
	CityBuster       bool
	AirAttacker      bool
	Horse            bool
	Submarine        bool
	BadCityDefender  bool
	OnlyNativeAttack bool
}

type DefenderCity struct {
	Size              int  `json:""`
	HasWalls          bool `json:""`
	HasCoastalDefence bool `json:""`
	HasSAM            bool `json:""`
	HasSDI            bool `json:""`
}

type DefenderTerrain struct {
	Type     string `json:""`
	HasRiver bool   `json:"hasRiver"`
}

type DefenderUnit struct {
	Name         string `json:"name"`
	VetLevel     int    `json:"vetLevel"`
	HP           int    `json:"hp"`
	HasRiver     bool   `json:"hasRiver"`
	HasCity      bool   `json:"hasCity"`
	HasFortress  bool   `json:"hasFortress"`
	IsFortified  bool   `json:"isFortified"`
	HasGreatWall bool   `json:""`

	City    DefenderCity    `json:"city"`
	Terrain DefenderTerrain `json:"terrain"`
}

type AttackerUnit struct {
	Name     string `json:"name"`
	VetLevel int    `json:"vetLevel"`
	HP       int    `json:"hp"`
}

type AttackerVDefender struct {
	Attacker AttackerUnit `json:"attacker"`
	Defender DefenderUnit `json:"defender"`
}
