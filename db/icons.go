package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

//IconID holds all Iconinformation from eve
type IconID struct {
	Backgrounds []string `json:"backgrounds,omitempty" yaml:"backgrounds"`
	Description string   `json:"description,omitempty" yaml:"description"`
	Foregrounds []string `json:"foregrounds,omitempty" yaml:"foregrounds"`
	IconFile    string   `json:"iconFile,omitempty" yaml:"iconFile"`
	Obsolete    bool     `json:"obsolete,omitempty" yaml:"obsolete"`
}

type iconIDs map[int]*IconID

func (tt iconIDs) GetByKey(key int) interface{} {
	return tt[key]
}

func (tt iconIDs) FileName() string {
	return "iconIDs"
}
func (tt iconIDs) SaveToDB() error {
	for k, v := range tt {
		if err := insert(tt.FileName(), k, v); err != nil {
			return err
		}
	}
	return nil
}
func (tt iconIDs) TableName() string {
	return strings.ToLower(tt.FileName())
}
func (tt iconIDs) New(id int, data []byte) error {
	var (
		err     error
		newItem = new(IconID)
	)
	if err = json.Unmarshal(data, newItem); err != nil {
		return err
	}
	tt[id] = newItem
	return nil
}
func (tt iconIDs) LoadFromDB() error {
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
