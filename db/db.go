package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	_ "github.com/lib/pq" //postgres
)

//EveDB is the entirety of the eve database tables
//implemented as maps of database id to actual type
type EveDB struct {
	Blueprints   map[int]*Blueprint
	CategoryIDs  map[int]*CategoryID
	Certificates map[int]*Certificate
	TypeIDs      map[int]*TypeID
	GraphicIDs   map[int]*GraphicID
	GroupIDs     map[int]*GroupID
	IconIDs      map[int]*IconID
	Skins        map[int]*Skin
	Mapping      map[string]interface{}
}

func (e *EveDB) keys(fileName string) []int {
	keys := make([]int, 0)
	switch fileName {
	case "blueprints":
		for k := range eve.Blueprints {
			keys = append(keys, k)
		}
	case "categoryIDs":
		for k := range eve.CategoryIDs {
			keys = append(keys, k)
		}
	case "certificates":
		for k := range eve.Certificates {
			keys = append(keys, k)
		}
	case "typeIDs":
		for k := range eve.TypeIDs {
			keys = append(keys, k)
		}
	case "graphicIDs":
		for k := range eve.GraphicIDs {
			keys = append(keys, k)
		}
	case "groupIDs":
		for k := range eve.GroupIDs {
			keys = append(keys, k)
		}
	case "iconIDs":
		for k := range eve.IconIDs {
			keys = append(keys, k)
		}
	case "skins":
		for k := range eve.Skins {
			keys = append(keys, k)
		}
	}
	return keys
}
func (e *EveDB)insertJSONByFileAndKey(fileName string, key int, data []byte)error{
	var (
		err error
	)
	switch fileName {
	case "blueprints":
		o := new(Blueprint)
		if err = json.Unmarshal(data,o);err!= nil{
			return err
		}
		eve.Blueprints[key]=o
		return nil
	case "categoryIDs":
		o := new(CategoryID)
		if err = json.Unmarshal(data,o);err!= nil{
			return err
		}
		eve.CategoryIDs[key]=o
		return nil
	case "certificates":
		o := new(Certificate)
		if err = json.Unmarshal(data,o);err!= nil{
			return err
		}
		eve.Certificates[key]=o
		return nil
	case "typeIDs":
		o := new(TypeID)
		if err = json.Unmarshal(data,o);err!= nil{
			return err
		}
		eve.TypeIDs[key]=o
		return nil
	case "graphicIDs":
		o := new(GraphicID)
		if err = json.Unmarshal(data,o);err!= nil{
			return err
		}
		eve.GraphicIDs[key]=o
		return nil
	case "groupIDs":
		o := new(GroupID)
		if err = json.Unmarshal(data,o);err!= nil{
			return err
		}
		eve.GroupIDs[key]=o
		return nil
	case "iconIDs":
		o := new(IconID)
		if err = json.Unmarshal(data,o);err!= nil{
			return err
		}
		eve.IconIDs[key]=o
		return nil
	case "skins":
		o := new(Skin)
		if err = json.Unmarshal(data,o);err!= nil{
			return err
		}
		eve.Skins[key]=o
		return nil
	}
	return err
}
func (e *EveDB) byFileName(fileName string, key int) (interface{}, bool) {
	switch fileName {
	case "blueprints":
		v, ok := eve.Blueprints[key]
		return v, ok
	case "categoryIDs":
		v, ok := eve.CategoryIDs[key]
		return v, ok
	case "certificates":
		v, ok := eve.Certificates[key]
		return v, ok
	case "typeIDs":
		v, ok := eve.TypeIDs[key]
		return v, ok
	case "graphicIDs":
		v, ok := eve.GraphicIDs[key]
		return v, ok
	case "groupIDs":
		v, ok := eve.GroupIDs[key]
		return v, ok
	case "iconIDs":
		v, ok := eve.IconIDs[key]
		return v, ok
	case "skins":
		v, ok := eve.Skins[key]
		return v, ok
	}
	return nil, false
}

//Keys retrieves all primary Keys from a table
//useful for inserts to only insert rows that
//are non existant
func Keys(table string) map[int]bool {
	var (
		rows *sql.Rows
		err  error
		key  int
		keys = make(map[int]bool)
	)
	if rows, err = pg.Query(`SELECT id FROM ` + table); err != nil {
		log.Fatalf("Can't Read Keys: %s", err)
	}
	for rows.Next() {
		if err = rows.Scan(&key); err != nil {
			log.Fatal(err)
		}
		keys[key] = true
	}
	return keys
}

//Init opens the database and stores it as a private
//package variable
func open() error {
	var err error
	pg, err = sql.Open("postgres", connectionString)
	return err

}

//Close the database all errors will make
//the Program crash so we don't have
//to deal with errors on defer
func Close() {
	if err := pg.Close(); err != nil {
		log.Fatalf("Can't Close Database: %s", err)
	}
}

//CreateIfNotExists creates a table as a key jsonb store
func createIfNotExists(table string) {
	if _, err := pg.Exec(`CREATE TABLE IF NOT EXISTS ` + table + ` (id integer PRIMARY KEY, data jsonb NOT NULL)`); err != nil {
		log.Fatalf("Can't create Table %s: %s", table, err)
	}
}

//doesTableExist returns wether the table exists
func tableExists(table string) bool {
	var exists = new(sql.NullString)
	row := pg.QueryRow(fmt.Sprintf("SELECT to_regclass('public.%s')", table))
	if err := row.Scan(exists); err != nil {
		log.Fatalf("Can't check if Table exists: %s", err)
	}
	return exists.Valid
}

//Insert a datastructure as JSONB into the database
func Insert(table string, key int, data interface{}) error {
	var (
		err      error
		jsonData []byte
	)
	if jsonData, err = json.Marshal(data); err != nil {
		return err
	}
	_, err = pg.Exec(`INSERT INTO `+table+`(id, data) VALUES($1, $2)`, key, jsonData)
	return err
}

func loadTable(table string){
	var (
		rows *sql.Rows
		err  error
	)
	rows, err = pg.Query(fmt.Sprintf("SELECT * FROM %s", table))
	for rows.Next() {
		id:= -1
		data := []byte("")
		if err = rows.Scan(&id, &data);err!= nil{
			log.Fatalf("Can't read value from %s: %s",table,err)
		}

		if err = eve.insertJSONByFileAndKey(table,id,data);err != nil{
			log.Fatalf("Couldn't insert ID:%v into table %s",id, table)
		}
	}
}
