package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/shaban/polaris/db"
	"github.com/spf13/viper"

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
		//typeIDs = make(map[int]*db.TypeID)
		//blueprints = make(map[int]*db.Blueprint)
	)
	defer db.Close()
	path := viper.GetString("path")
	yamlPath := viper.GetString("database.yaml.path")
	ext:= viper.GetString("database.yaml.extension")
	paths := viper.GetStringSlice("database.yaml.files")

	eve := new(db.EveDB)

	for _, p := range paths{
		//tableMap := make(map[int]interface{})
		filePath := fmt.Sprintf("%s/%s/%s.%s",path, yamlPath,p,ext)
		switch p{
		case "blueprints":
			if err = db.LoadYaml(&eve.Blueprints,filePath);err!= nil{
				log.Fatal(err)
			}
		case "categoryIDs":
			if err = db.LoadYaml(&eve.CategoryIDs,filePath);err!= nil{
				log.Fatal(err)
			}
		case "certificates":
			if err = db.LoadYaml(&eve.Certificates,filePath);err!= nil{
				log.Fatal(err)
			}
		case "typeIDs":
			if err = db.LoadYaml(&eve.TypeIDs,filePath);err!= nil{
				log.Fatal(err)
			}
			return
		}
		//db.CreateIfNotExists(p)
		println(filePath)
	}
	for k, v := range eve.Blueprints{
		if v.Activities.Reaction.Time != 0{
			t, _ := eve.TypeIDs[k]
			println(t.Name.En, v.Activities.Reaction.Time)
		}
		
	}
	//categoryIDs := viper.GetString("database.yaml.categoryIDs")
	/*if err = db.LoadYaml(&typeIDs, path+categoryIDs); err != nil {
		log.Fatal(err.Error())
	}
	keys := db.Keys("esi")
	for k, v := range typeIDs {
		if _, isInTable := keys[k]; !isInTable {
			db.Insert("esi", k, v)
		}
	}*/
}
