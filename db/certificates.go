package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
)

//CertificateLevel holds the ranks needed
//to fulfill a certificates respective level
type CertificateLevel struct {
	Advanced int `json:"advanced,omitempty" yaml:"advanced"`
	Basic    int `json:"basic,omitempty" yaml:"basic"`
	Elite    int `json:"elite,omitempty" yaml:"elite"`
	Improved int `json:"improved,omitempty" yaml:"improved"`
	Standard int `json:"standard,omitempty" yaml:"standard"`
}

//Certificate holds all information
//on a given eve certificate
type Certificate struct {
	Description    string                    `json:"description,omitempty" yaml:"description"`
	GroupID        int                       `json:"groupID,omitempty" yaml:"groupID"`
	Name           string                    `json:"name,omitempty" yaml:"name"`
	RecommendedFor []int                     `json:"recommendedFor,omitempty" yaml:"recommendedFor"`
	SkillTypes     map[int]*CertificateLevel `json:"skillTypes,omitempty" yaml:"skillTypes"`
}

type certificates map[int]*Certificate

func (tt certificates) GetByKey(key int) interface{} {
	return tt[key]
}

func (tt certificates) FileName() string {
	return "certificates"
}
func (tt certificates) SaveToDB() error {
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
func (tt certificates) TableName() string {
	return strings.ToLower(tt.FileName())
}
func (tt certificates) New(id int, data []byte) error {
	var (
		err     error
		newItem = new(Certificate)
	)
	if err = json.Unmarshal(data, newItem); err != nil {
		return err
	}
	tt[id] = newItem
	return nil
}
func (tt certificates) LoadFromDB() error {
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
