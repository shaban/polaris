package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq" //postgres
	"github.com/spf13/viper"
)

var (
	pg *sql.DB
)

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

//Open the database and store it as private
//package variable
func Init() error {
	var err error
	connStr := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s",
		viper.GetString("database.postgres.user"),
		viper.GetString("database.postgres.db"),
		viper.GetString("database.postgres.password"),
		viper.GetString("database.postgres.host"),
	)
	pg, err = sql.Open("postgres", connStr)
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
func CreateIfNotExists(table string){
	if _, err := pg.Exec(`CREATE TABLE IF NOT EXISTS `+table+` (id integer PRIMARY KEY, data jsonb NOT NULL)`);err != nil{
		log.Fatalf("Can't create Table %s: %s",table,err)
	}
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
