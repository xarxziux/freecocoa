package models

// Enums
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

// Input stats: the data structures in the original requests
type AttackerInput struct {
	Name     string `json:"name"`
	VetLevel int    `json:"vetLevel"`
	HP       int    `json:"hp"`
	MP       int    `json:"mp"`
}

type DefenderInput struct {
	Name         string `json:"name"`
	VetLevel     int    `json:"vetLevel"`
	HP           int    `json:"hp"`
	IsFortified  bool   `json:"isFortified"`
	HasFortress  bool   `json:"hasFortress"`
	HasAirbase   bool   `json:"isInAirbase"`
	HasCity      bool   `json:"hasCity"`
	HasGreatWall bool   `json:"hasGreatWall"`
}

type CityInput struct {
	Size              int  `json:"size"`
	HasWalls          bool `json:"hasWalls"`
	HasCoastalDefense bool `json:"hasCoastalDefense"`
	HasSAM            bool `json:"hasSAM"`
	SDILevel          int  `json:"sdiLevel"`
	IsCapital         bool `json:"isCapital"`
}

type TerrainInput struct {
	Type     string `json:"type"`
	HasRiver bool   `json:"hasRiver"`
}

// Endpoint stats: full data structures used by the various endpoints
type AttackInput struct {
	Attacker AttackerInput `json:"attacker"`
	Defender DefenderInput `json:"defender"`
	City     CityInput     `json:"city"`
	Terrain  TerrainInput  `json:"terrain"`
}

type AttackAllInput struct {
	Attackers []AttackerInput `json:"attackers"`
	Defenders []DefenderInput `json:"defenders"`
	City      CityInput       `json:"city"`
	Terrain   TerrainInput    `json:"terrain"`
}

// Base stats: the basic, modifiable, unit-specific stats when all factors have been calculated
type AttackerBase struct {
	AP float64 `json:"attackPower"`
	HP int     `json:"hitpoints"`
	FP int     `json:"firepower"`
	MP int     `json:"movementPoints,omitempty"`
}

type DefenderBase struct {
	DP float64 `json:"defensePower"`
	HP int     `json:"hitpoints"`
	FP int     `json:"firepower"`
	MP int     `json:"movementPoints,omitempty"`
}

// Detail stats: fixed stats shared by all units and terrain of the same type
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
	MP               int
	BuildCost        int
	CityBuster       bool
	AirAttacker      bool
	Horse            bool
	Submarine        bool
	BadCityDefender  bool
	OnlyNativeAttack bool
	CantFortify      bool
}

type TerrainDetails struct {
	Name         string
	Class        TerrainClass
	DefenseBonus int
	CanHaveRiver bool
	NoCities     bool
	NoFortify    bool
	UnsafeCoast  bool
}

// Pair stats: combine fixed and modifiable stats
type AttackerPair struct {
	Base    AttackerBase
	Details UnitDetails
}

type DefenderPair struct {
	Base    DefenderBase
	Details UnitDetails
}

// Validated stats
type AttackerValidated struct {
	Input   AttackerInput
	Details UnitDetails
}

type DefenderValidated struct {
	Input   DefenderInput
	Details UnitDetails
}

type TerrainValidated struct {
	Input   TerrainInput
	Details TerrainDetails
}

type AttackValidated struct {
	Attacker AttackerValidated `json:"attacker"`
	Defender DefenderValidated `json:"defender"`
	City     CityInput         `json:"city"`
	Terrain  TerrainValidated  `json:"terrain"`
}

type AttackAllValidated struct {
	Attackers []*AttackerValidated `json:"attacker"`
	Defenders []*DefenderValidated `json:"defender"`
	City      CityInput            `json:"city"`
	Terrain   TerrainValidated     `json:"terrain"`
}

// Results stats: the final values returned by the endpoints
type AttackResults struct {
	Attacker AttackerBase `json:"attacker"`
	Defender DefenderBase `json:"defender"`
}

type CombatResult struct {
	AttProb  string `json:"attackProbability"`
	AttHP    string `json:"attackerHP"`
	DefProb  string `json:"defenderProbability"`
	DefHP    string `json:"defenderHP"`
	DefDelta string `json:"defenderDelta"`
}

type CombatResults struct {
	Prob    float64         `json:"probability"`
	Results []*CombatResult `json:"results"`
}

type CombinedResults struct {
	Stats  *AttackResults `json:"stats"`
	Combat *CombatResults `json:"combat"`
}

// Data types for the buildCost endpoint
type BuildCostInput struct {
	UnitName      string `json:"unitName"`
	ShieldCost    int    `json:"shieldCost"`
	ShieldOutput  int    `json:"shieldOutput"`
	ShieldCurrent int    `json:"shieldCurrent"`
}

type BuildCostItem struct {
	ShieldCurrent   int `json:"current"`
	ShieldRemaining int `json:"remaining"`
	BuyCost         int `json:"buyCost"`
}

type BuildCostOutput struct {
	BuildCosts []BuildCostItem `json:"costs"`
}

var Exists = struct{}{}
