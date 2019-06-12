package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

type MarketGroup struct {
	Description     string `json:"description,omitempty" yaml:"description"`
	HasTypes        bool   `json:"hasTypes,omitempty" yaml:"hasTypes"`
	IconID          int    `json:"iconID,omitempty" yaml:"iconID"`
	MarketGroupID   int    `json:"marketGroupID,omitempty" yaml:"marketGroupID"`
	MarketGroupName string `json:"marketGroupName,omitempty" yaml:"marketGroupName"`
	ParentGroupID   int    `json:"parentGroupID,omitempty" yaml:"parentGroupID"`
}

type marketGroups map[int]*MarketGroup

func (tt marketGroups) GetByKey(key int) interface{} {
	return tt[key]
}
func (tt marketGroups) SaveToDB() error {
	for k, v := range tt {
		if err := insert(tt.FileName(), k, v); err != nil {
			return err
		}
	}
	return nil
}
func (tt marketGroups) FileName() string {
	return "invMarketGroups"
}
func (tt marketGroups) TableName() string {
	return strings.ToLower(tt.FileName())
}
func (tt marketGroups) New(id int, data []byte) error {
	var (
		err     error
		newItem = new(MarketGroup)
	)
	if err = json.Unmarshal(data, newItem); err != nil {
		return fmt.Errorf("Can't load into Table:%s ID:%v %s", tt.TableName(), id, data)
	}
	tt[id] = newItem
	return nil
}
func (tt marketGroups) LoadFromDB() error {
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
func (tt marketGroups) LoadFromYAML(dec *yaml.Decoder) error {
	var (
		err error
		arr = make([]*MarketGroup, 0)
	)
	if err = dec.Decode(&arr); err != nil {
		return err
	}
	for _, v := range arr {
		tt[v.MarketGroupID] = v
	}
	return nil
}
