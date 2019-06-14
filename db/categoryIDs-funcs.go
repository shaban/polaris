// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package db

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type categoryIDs map[int]*CategoryID

func (tt categoryIDs) GetByKey(key int) interface{} {
	return tt[key]
}
func (tt categoryIDs) SaveToDB() error {
	for k, v := range tt {
		if err := insert(tt.FileName(), k, v); err != nil {
			return err
		}
	}
	return nil
}
func (tt categoryIDs) FileName() string {
	return "categoryIDs"
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
		return fmt.Errorf("Can't load into Table:%s ID:%v %s", tt.TableName(), id, data)
	}
	tt[id] = newItem
	return nil
}
func (tt categoryIDs) LoadFromDB() error {
	return loadFromDB(tt)
}
func (tt categoryIDs) LoadFromYAML() error {
	path := fmt.Sprintf("%s/%s/%s.%s", basePath, yamlObjectPath, tt.FileName(), yamlExt)
	var (
		err error
		f   *os.File
	)
	if f, err = os.OpenFile(path, os.O_RDONLY, 0644); err != nil {
		return err
	}
	defer f.Close()
	dec := yaml.NewDecoder(f)
	dec.SetStrict(true)
	return dec.Decode(&tt)
}
