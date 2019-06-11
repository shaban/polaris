package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

//GroupID holds all Item Groups of Eve
type GroupID struct {
	Anchorable           bool       `json:"anchorable,omitempty" yaml:"anchorable"`
	Anchored             bool       `json:"anchored,omitempty" yaml:"anchored"`
	CategoryID           int        `json:"categoryID,omitempty" yaml:"categoryID"`
	FittableNonSingleton bool       `json:"fittableNonSingleton,omitempty" yaml:"fittableNonSingleton"`
	Name                 translated `json:"name,omitempty" yaml:"name"`
	Published            bool       `json:"published,omitempty" yaml:"published"`
	UseBasePrice         bool       `json:"useBasePrice,omitempty" yaml:"useBasePrice"`
	IconID               int        `json:"iconID,omitempty" yaml:"iconID"`
}

type groupIDs map[int]*GroupID

func (tt groupIDs) GetByKey(key int) interface{} {
	return tt[key]
}

func (tt groupIDs) FileName() string {
	return "groupIDs"
}
func (tt groupIDs) SaveToDB() error {
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
func (tt groupIDs) TableName() string {
	return strings.ToLower(tt.FileName())
}
func (tt groupIDs) New(id int, data []byte) error {
	var (
		err     error
		newItem = new(GroupID)
	)
	if err = json.Unmarshal(data, newItem); err != nil {
		return err
	}
	tt[id] = newItem
	return nil
}
func (tt groupIDs) LoadFromDB() error {
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
