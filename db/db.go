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
const(
	CategoryIDs = "database.yaml.categoryIDs"
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
func Close() {
	if err := pg.Close(); err != nil {
		log.Fatalf("Can't Close Database: %s", err)
	}
}

func Insert(table string, key int, data interface{}) error {
	var (
		err      error
		jsonData []byte
	)
	if jsonData, err = json.Marshal(data); err != nil {
		return err
	}
	_, err = pg.Exec(`INSERT INTO esi(id, data) VALUES($1, $2)`, key, jsonData)
	return err
}
