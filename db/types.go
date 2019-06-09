package db

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