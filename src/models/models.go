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
	Amphibious
	AncientLand
	Air
	BigLand
	BigSiege
	DeepSea
	Helicopter
	Land
	Merchant
	Missile
	Nuclear
	Patrol
	Sea
	SmallLand
	SmallSea
	SmallUnit
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

type TerrainType struct {
	Name         string
	Class        TerrainClass
	DefenseBonus uint8
	CanHaveRiver bool
	NoCities     bool
	NoFortify    bool
	UnsafeCoast  bool
}

type CityStats struct {
	Size           uint8 `json:"size"`
	Walls          bool  `json:"walls"`
	GreatWalls     bool  `json:"greatWalls"`
	CoastalDefence bool  `json:"coastalDefence"`
}

type AttackStats struct {
	AP float32 `json:"attack_power"`
	HP uint8   `json:"hitpoints"`
	FP uint8   `json:"firepower"`
}

type DefenceStats struct {
	DP float32 `json:"defense_power"`
	HP uint8   `json:"hitpoints"`
	FP uint8   `json:"firepower"`
}

type FinalStats struct {
	Attacker AttackStats  `json:"attacker"`
	Defender DefenceStats `json:"defender"`
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
	AP               float32
	DP               float32
	HP               uint8
	FP               uint8
	CityBuster       bool
	AirAttacker      bool
	Horse            bool
	Submarine        bool
	BadCityDefender  bool
	OnlyNativeAttack bool
	CantFortify      bool
}

type PairDetails struct {
	Attacker UnitDetails
	Defender UnitDetails
}

type DefenderCity struct {
	Size              uint8 `json:"size"`
	HasWalls          bool  `json:"hasWalls"`
	HasGreatWall      bool  `json:"hasGreatWall"`
	HasCoastalDefence bool  `json:"hasCoastalDefence"`
	HasSAM            bool  `json:"hasSAM"`
	SDILevel          uint8 `json:"sdiLevel"`
	IsCapital         bool  `json:"isCapital"`
}

type DefenderTerrain struct {
	Type     string `json:"type"`
	HasRiver bool   `json:"hasRiver"`
}

type DefenderUnit struct {
	Name         string `json:"name"`
	VetLevel     uint8  `json:"vetLevel"`
	HP           uint8  `json:"hp"`
	HasCity      bool   `json:"hasCity"`
	HasFortress  bool   `json:"hasFortress"`
	IsFortified  bool   `json:"isFortified"`
	HasGreatWall bool   `json:""`

	City    DefenderCity    `json:"city"`
	Terrain DefenderTerrain `json:"terrain"`
}

type AttackerUnit struct {
	Name     string `json:"name"`
	VetLevel uint8  `json:"vetLevel"`
	HP       uint8  `json:"hp"`
}

type AttackerVDefender struct {
	Attacker AttackerUnit `json:"attacker"`
	Defender DefenderUnit `json:"defender"`
}

type BaseStats struct {
	Input   AttackerVDefender
	Details PairDetails
	Terrain TerrainType
}
