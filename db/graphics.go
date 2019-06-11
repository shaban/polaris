package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

//GraphicID holds the graphic descriptions from Eve
type GraphicID struct {
	Description    string   `json:"description,omitempty" yaml:"description"`
	GraphicFile    string   `json:"graphicFile,omitempty" yaml:"graphicFile"`
	IconInfo       iconInfo `json:"iconInfo,omitempty" yaml:"iconInfo"`
	SofFactionName string   `json:"sofFactionName,omitempty" yaml:"sofFactionName"`
	SofHullName    string   `json:"sofHullName,omitempty" yaml:"sofHullName"`
	SofRaceName    string   `json:"sofRaceName,omitempty" yaml:"sofRaceName"`
}
type iconInfo struct {
	Backgrounds []string `json:"backgrounds,omitempty" yaml:"backgrounds"`
	Folder      string   `json:"folder,omitempty" yaml:"folder"`
	Foregrounds []string `json:"foregrounds,omitempty" yaml:"foregrounds"`
}

type graphicIDs map[int]*GraphicID

func (tt graphicIDs) GetByKey(key int) interface{} {
	return tt[key]
}
func (tt graphicIDs) FileName() string {
	return "graphicIDs"
}
func (tt graphicIDs) SaveToDB() error {
	for k, v := range tt {
		if err := insert(tt.FileName(), k, v); err != nil {
			return err
		}
	}
	return nil
}
func (tt graphicIDs) TableName() string {
	return strings.ToLower(tt.FileName())
}
func (tt graphicIDs) New(id int, data []byte) error {
	var (
		err     error
		newItem = new(GraphicID)
	)
	if err = json.Unmarshal(data, newItem); err != nil {
		return err
	}
	tt[id] = newItem
	return nil
}
func (tt graphicIDs) LoadFromDB() error {
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
