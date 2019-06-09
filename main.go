package main

import (
	//"bytes"
	//"database/sql"
	//"encoding/json"

	_ "github.com/lib/pq"
	"github.com/shaban/polaris/db"
	"github.com/spf13/viper"

	//"io/ioutil"
	//"fmt"
	"log"
)

func init() {
	var (
		err error
	)
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/code/src/github.com/shaban/polaris")
	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	if err = db.Init(); err != nil {
		log.Fatal(err)
	}
}
func main() {
	var (
		err     error
		typeIDs = make(map[int]*db.TypeID)
		//rows    *sql.Rows
		//key     int
		//keys    map[int]bool
	)
	defer db.Close()
	path := viper.GetString("path")
	categoryIDs := viper.GetString("database.yaml.categoryIDs")
	if err = db.LoadYaml(&typeIDs, path+categoryIDs); err != nil {
		log.Fatal(err.Error())
	}
	keys := db.Keys("esi")
	for k, v := range typeIDs{
		if _, isInTable := keys[k];!isInTable{
			db.Insert("esi",k,v)
		}
	}
}
