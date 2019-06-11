package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"

	_ "github.com/lib/pq" //postgres
)

//EveDB is the entirety of the eve database tables
//implemented as maps of database id to actual type
type EveDB struct {
	Mapping      map[string]mapper
}

type mapper interface {
	GetByKey(int) interface{}
	SaveToDB() error
	FileName() string
	LoadFromDB() error
	TableName() string
	New(int, []byte) error
}
//EncodeByTableAndKey accesses items by tablename, row id
//and then encodes the row in JSON into a writer interface
func EncodeByTableAndKey(w io.Writer, table string, key int)error{
	var (
		v interface{}

	)
	if v = eve.Mapping[table].GetByKey(key); v == nil {
		return fmt.Errorf("Can't get table:%s key:%v is nil", table, key)
	}
	enc := json.NewEncoder(w)
	return enc.Encode(v)
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
func insert(table string, key int, data interface{}) error {
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
