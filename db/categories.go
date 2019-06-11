package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

//CategoryID hols a categoryID
type CategoryID struct {
	IconID    int        `json:"iconID,omitempty" yaml:"iconID"`
	Name      translated `json:"name,omitempty" yaml:"name"`
	Published bool       `json:"published,omitempty" yaml:"published"`
}

type categoryIDs map[int]*CategoryID

func (tt categoryIDs) GetByKey(key int) interface{} {
	return tt[key]
}
func (tt categoryIDs) FileName() string {
	return "categoryIDs"
}
func (tt categoryIDs) SaveToDB() error {
	for k, v := range tt {
		if err := insert(tt.FileName(), k, v); err != nil {
			return err
		}
	}
	return nil
}

func (tt categoryIDs) TableName() string {
	return strings.ToLower(tt.FileName())
}
func (tt categoryIDs) New(id int, data []byte) error {
	var (
		err     error
		newItem = new(CategoryID)
	)
	if err = json.Unmarshal(data, newItem); err != nil {
		return err
	}
	tt[id] = newItem
	return nil
}
func (tt categoryIDs) LoadFromDB() error {
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
