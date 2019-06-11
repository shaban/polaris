package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

//Skin holds ship skin informations
type Skin struct {
	AllowCCPDevs       bool   `json:"allowCCPDevs,omitempty" yaml:"allowCCPDevs"`
	InternalName       string `json:"internalName,omitempty" yaml:"internalName"`
	SkinID             int    `json:"skinID,omitempty" yaml:"skinID"`
	SkinMaterialID     int    `json:"skinMaterialID,omitempty" yaml:"skinMaterialID"`
	Types              []int  `json:"types,omitempty" yaml:"types"`
	VisibleSerenity    bool   `json:"visibleSerenity,omitempty" yaml:"visibleSerenity"`
	VisibleTranquility bool   `json:"visibleTranquility,omitempty" yaml:"visibleTranquility"`
	SkinDescription    string `json:"skinDescription,omitempty" yaml:"skinDescription"`
	IsStructureSkin    bool   `json:"isStructureSkin,omitempty" yaml:"isStructureSkin"`
}
type skins map[int]*Skin

func (tt skins) GetByKey(key int) interface{} {
	return tt[key]
}
func (tt skins) FileName() string {
	return "skins"
}
func (tt skins) SaveToDB() error {
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
func (tt skins) TableName() string {
	return strings.ToLower(tt.FileName())
}
func (tt skins) New(id int, data []byte) error {
	var (
		err     error
		newItem = new(Skin)
	)
	if err = json.Unmarshal(data, newItem); err != nil {
		return err
	}
	tt[id] = newItem
	return nil
}
func (tt skins) LoadFromDB() error {
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
