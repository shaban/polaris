package db

type EveDB struct {
	Blueprints   map[int]*Blueprint
	CategoryIDs  map[int]*CategoryID
	Certificates map[int]*Certificate
	TypeIDs      map[int]*TypeID
}

//Materials is a Blueprint Material
type Materials struct {
	Quantity int `json:"quantity"`
	TypeID   int `json:"typeID"`
}

//Skills is a Blueprint Process Requirement
type Skills struct {
	Level  int `json:"level"`
	TypeID int `json:"typeID"`
}

//Process is a Blueprint Activities Process
type Process struct {
	Materials []Materials `json:"materials"`
	Products  []Materials `json:"products"`
	Skills    []Skills    `json:"skills"`
	Time      int         `json:"time"`
}

//Activities is a Blueprint Activity
type Activities struct {
	Reaction         Process `json:"reaction"`
	Copying          Process `json:"copying"`
	Invention        Process `json:"invention"`
	Manufacturing    Process `json:"manufacturing"`
	ResearchMaterial Process `json:"research_material"`
	ResearchTime     Process `json:"research_time"`
}

//Blueprint holds blueprint information
type Blueprint struct {
	Activities         `json:"activities"`
	BlueprintTypeID    int `json:"blueprintTypeID"`
	MaxProductionLimit int `json:"maxProductionLimit"`
}

//CategoryID hols a categoryID
type CategoryID struct {
	IconID int `json:"iconID"`
	Name   struct {
		De string `json:"de"`
		En string `json:"en"`
		Fr string `json:"fr"`
		Ja string `json:"ja"`
		Ru string `json:"ru"`
		Zh string `json:"zh"`
	} `json:"name"`
	Published bool `json:"published"`
}
type CertificateLevel struct {
	Advanced int `json:"advanced"`
	Basic    int `json:"basic"`
	Elite    int `json:"elite"`
	Improved int `json:"improved"`
	Standard int `json:"standard"`
}
type Certificate struct {
	Description    string                    `json:"description"`
	GroupID        int                       `json:"groupID"`
	Name           string                    `json:"name"`
	RecommendedFor []int                     `json:"recommendedFor"`
	SkillTypes     map[int]*CertificateLevel `json:"skillTypes"`
}

//TypeID holds the item information
type TypeID struct {
	GroupID       int               `yaml:"groupID"`
	Name          map[string]string `yaml:"name"`
	Description   map[string]string `yaml:"description"`
	Mass          float64           `yaml:"mass"`
	Volume        float64           `yaml:"volume"`
	Capacity      float64           `yaml:"capacity"`
	PortionSize   int               `yaml:"portionSize"`
	RaceID        int               `yaml:"raceID"`
	BasePrice     float64           `yaml:"basePrice"`
	Published     bool              `yaml:"published"`
	MarketGroupID int               `yaml:"marketGroupID"`
	Radius        float64           `yaml:"radius"`
	IconID        int               `yaml:"iconID"`
	SoundID       int               `yaml:"soundID"`
	GraphicID     int               `yaml:"graphicID"`
	FactionName   string            `yaml:"sofFactionName"`
	FactionID     int               `yaml:"factionID"`
	MaterialSetID int               `yaml:"sofMaterialSetID"`
	Masteries     map[int][]int     `yaml:"masteries"`
	*Trait        `yaml:"traits"`
}

//Trait holds all the traits
type Trait struct {
	Types       map[int][]*RoleBonus `yaml:"types"`
	RoleBonuses []*RoleBonus         `yaml:"roleBonuses"`
	MiscBonuses []*RoleBonus         `yaml:"miscBonuses"`
}

//RoleBonus is a ship/subsystem trait
type RoleBonus struct {
	Bonus      int               `yaml:"bonus"`
	BonusText  map[string]string `yaml:"bonusText"`
	Importance int               `yaml:"importance"`
	UnitID     int               `yaml:"unitID"`
}
