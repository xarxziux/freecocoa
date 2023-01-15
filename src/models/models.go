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
	DefenseBonus int
	CanHaveRiver bool
	NoCities     bool
	NoFortify    bool
	UnsafeCoast  bool
}

type AttackStats struct {
	AP float64 `json:"attackPower"`
	HP int     `json:"hitpoints"`
	FP int     `json:"firepower"`
}

type DefenceStats struct {
	DP float64 `json:"defensePower"`
	HP int     `json:"hitpoints"`
	FP int     `json:"firepower"`
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
	AP               float64
	DP               float64
	HP               int
	FP               int
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
	Size              int  `json:"size"`
	HasWalls          bool `json:"hasWalls"`
	HasCoastalDefence bool `json:"hasCoastalDefence"`
	HasSAM            bool `json:"hasSAM"`
	SDILevel          int  `json:"sdiLevel"`
	IsCapital         bool `json:"isCapital"`
}

type DefenderTerrain struct {
	Type     string `json:"type"`
	HasRiver bool   `json:"hasRiver"`
}

type DefenderUnit struct {
	Name         string `json:"name"`
	VetLevel     int    `json:"vetLevel"`
	HP           int    `json:"hp"`
	HasCity      bool   `json:"hasCity"`
	HasFortress  bool   `json:"hasFortress"`
	IsFortified  bool   `json:"isFortified"`
	HasGreatWall bool   `json:"hasGreatWall"`

	City    DefenderCity    `json:"city"`
	Terrain DefenderTerrain `json:"terrain"`
}

type AttackerUnit struct {
	Name     string `json:"name"`
	VetLevel int    `json:"vetLevel"`
	HP       int    `json:"hp"`
	MP       int    `json:"mp"`
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

type CombatResult struct {
	AttProb  string `json:"attackProbability"`
	AttHP    string `json:"attackerHP"`
	DefProb  string `json:"defenderProbability"`
	DefHP    string `json:"defenderHP"`
	DefDelta string `json:"defenderDelta"`
}

type CombatResults struct {
	Prob    float64        `json:"probability"`
	Results []CombatResult `json:"results"`
}

type CombinedResults struct {
	Stats  FinalStats    `json:"stats"`
	Combat CombatResults `json:"combat"`
}
