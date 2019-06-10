package db

//EveDB is the entirety of the eve database tables
//implemented as maps of database id to actual type
type EveDB struct {
	Blueprints   map[int]*Blueprint
	CategoryIDs  map[int]*CategoryID
	Certificates map[int]*Certificate
	TypeIDs      map[int]*TypeID
}

type translated struct {
	De string `json:"de" yaml:"de"`
	En string `json:"en" yaml:"en"`
	Es string `json:"es" yaml:"es"`
	Fr string `json:"fr" yaml:"fr"`
	It string `json:"it" yaml:"it"`
	Ja string `json:"ja" yaml:"ja"`
	Ru string `json:"ru" yaml:"ru"`
	Zh string `json:"zh" yaml:"zh"`
}

//CategoryID hols a categoryID
type CategoryID struct {
	IconID    int        `json:"iconID" yaml:"iconID"`
	Name      translated `json:"name" yaml:"name"`
	Published bool       `json:"published" yaml:"published"`
}

//TypeID holds the item information
type TypeID struct {
	GroupID       int           `json:"groupID" yaml:"groupID"`
	Name          translated    `json:"name" yaml:"name"`
	Description   translated    `json:"description" yaml:"description"`
	Mass          float64       `json:"mass" yaml:"mass"`
	Volume        float64       `json:"volume" yaml:"volume"`
	Capacity      float64       `json:"capacity" yaml:"capacity"`
	PortionSize   int           `json:"portionSize" yaml:"portionSize"`
	RaceID        int           `json:"raceID" yaml:"raceID"`
	BasePrice     float64       `json:"basePrice" yaml:"basePrice"`
	Published     bool          `json:"published" yaml:"published"`
	MarketGroupID int           `json:"marketGroupID" yaml:"marketGroupID"`
	Radius        float64       `json:"radius" yaml:"radius"`
	IconID        int           `json:"iconID" yaml:"iconID"`
	SoundID       int           `json:"soundID" yaml:"soundID"`
	GraphicID     int           `json:"graphicID" yaml:"graphicID"`
	FactionName   string        `json:"sofFactionName" yaml:"sofFactionName"`
	FactionID     int           `json:"factionID" yaml:"factionID"`
	MaterialSetID int           `json:"sofMaterialSetID" yaml:"sofMaterialSetID"`
	Masteries     map[int][]int `json:"masteries" yaml:"masteries"`
	Traits trait        `json:"traits" yaml:"traits"`
}

//Trait holds all the traits
type trait struct {
	Types       map[int][]*roleBonus `json:"types" yaml:"types"`
	RoleBonuses []*roleBonus         `json:"roleBonuses" yaml:"roleBonuses"`
	MiscBonuses []*roleBonus         `json:"miscBonuses" yaml:"miscBonuses"`
}

//RoleBonus is a ship/subsystem trait
type roleBonus struct {
	Bonus      int               `json:"bonus" yaml:"bonus"`
	BonusText  map[string]string `json:"bonusText" yaml:"bonusText"`
	Importance int               `json:"importance" yaml:"importance"`
	UnitID     int               `json:"unitID" yaml:"unitID"`
}
