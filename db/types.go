package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	_ "github.com/lib/pq" //postgres
)

type translated struct {
	De string `json:"de,omitempty" yaml:"de"`
	En string `json:"en,omitempty" yaml:"en"`
	Es string `json:"es,omitempty" yaml:"es"`
	Fr string `json:"fr,omitempty" yaml:"fr"`
	It string `json:"it,omitempty" yaml:"it"`
	Ja string `json:"ja,omitempty" yaml:"ja"`
	Ru string `json:"ru,omitempty" yaml:"ru"`
	Zh string `json:"zh,omitempty" yaml:"zh"`
}

//TypeID holds the item information
type TypeID struct {
	GroupID       int           `json:"groupID,omitempty" yaml:"groupID"`
	Name          translated    `json:"name,omitempty" yaml:"name"`
	Description   translated    `json:"description,omitempty" yaml:"description"`
	Mass          float64       `json:"mass,omitempty" yaml:"mass"`
	Volume        float64       `json:"volume,omitempty" yaml:"volume"`
	Capacity      float64       `json:"capacity,omitempty" yaml:"capacity"`
	PortionSize   int           `json:"portionSize,omitempty" yaml:"portionSize"`
	RaceID        int           `json:"raceID,omitempty" yaml:"raceID"`
	BasePrice     float64       `json:"basePrice,omitempty" yaml:"basePrice"`
	Published     bool          `json:"published,omitempty" yaml:"published"`
	MarketGroupID int           `json:"marketGroupID,omitempty" yaml:"marketGroupID"`
	Radius        float64       `json:"radius,omitempty" yaml:"radius"`
	IconID        int           `json:"iconID,omitempty" yaml:"iconID"`
	SoundID       int           `json:"soundID,omitempty" yaml:"soundID"`
	GraphicID     int           `json:"graphicID,omitempty" yaml:"graphicID"`
	FactionName   string        `json:"sofFactionName,omitempty" yaml:"sofFactionName"`
	FactionID     int           `json:"factionID,omitempty" yaml:"factionID"`
	MaterialSetID int           `json:"sofMaterialSetID,omitempty" yaml:"sofMaterialSetID"`
	Masteries     map[int][]int `json:"masteries,omitempty" yaml:"masteries"`
	Traits        trait         `json:"traits,omitempty" yaml:"traits"`
}

//Trait holds all the traits
type trait struct {
	Types       map[int][]*roleBonus `json:"types,omitempty" yaml:"types"`
	RoleBonuses []*roleBonus         `json:"roleBonuses,omitempty" yaml:"roleBonuses"`
	MiscBonuses []*roleBonus         `json:"miscBonuses,omitempty" yaml:"miscBonuses"`
}

//RoleBonus is a ship/subsystem trait
type roleBonus struct {
	Bonus      int               `json:"bonus,omitempty" yaml:"bonus"`
	BonusText  map[string]string `json:"bonusText,omitempty" yaml:"bonusText"`
	Importance int               `json:"importance,omitempty" yaml:"importance"`
	UnitID     int               `json:"unitID,omitempty" yaml:"unitID"`
}
type typeIDs map[int]*TypeID

func (tt typeIDs) GetByKey(key int) interface{} {
	return tt[key]
}

func (tt typeIDs) FileName() string {
	return "typeIDs"
}
func (tt typeIDs) SaveToDB() error {
	for k, v := range tt {
		b, err := json.Marshal(v)
		if err != nil {
			return err
		}
		if err = insert(tt.FileName(), k, b); err != nil {
			return err
		}
	}
	return nil
}
func (tt typeIDs) TableName() string {
	return strings.ToLower(tt.FileName())
}
func (tt typeIDs) New(id int, data []byte) error {
	var (
		err     error
		newItem = new(TypeID)
	)
	if err = json.Unmarshal(data, newItem); err != nil {
		return err
	}
	tt[id] = newItem
	return nil
}
func (tt typeIDs) LoadFromDB() error {
	var (
		rows *sql.Rows
		err  error
		id   int
		data []byte
	)
	rows, err = pg.Query(fmt.Sprintf("SELECT * FROM %s", tt.FileName()))
	for rows.Next() {
		if err = rows.Scan(&id, &data); err != nil {
			return fmt.Errorf("Can't read value from %s: %s", tt.FileName(), err)
		}

		if err = tt.New(id, data); err != nil {
			return err
		}
	}
	return nil
}
