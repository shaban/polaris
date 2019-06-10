package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var (
	connectionString string
	basePath         string
	yamlPath         string
	yamlExt          string
	yamlFiles        []string
	//fileAllocations  map[string]interface{}
	pg               *sql.DB
	eve              *EveDB
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
	connectionString = fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s",
		viper.GetString("database.postgres.user"),
		viper.GetString("database.postgres.db"),
		viper.GetString("database.postgres.password"),
		viper.GetString("database.postgres.host"),
	)

	if err = open(); err != nil {
		log.Fatal(err)
	}

	basePath = viper.GetString("path")
	yamlPath = viper.GetString("database.yaml.path")
	yamlExt = viper.GetString("database.yaml.extension")
	yamlFiles = viper.GetStringSlice("database.yaml.files")

	eve = new(EveDB)
	eve.Mapping = make(map[string]interface{})

	for _, fileName := range yamlFiles{
		switch fileName{
		case "blueprints":
			eve.Blueprints=make(map[int]*Blueprint)
			eve.Mapping[fileName]=&eve.Blueprints
		case "categoryIDs":
			eve.CategoryIDs=make(map[int]*CategoryID)
			eve.Mapping[fileName]=&eve.CategoryIDs
		case "certificates":
			eve.Certificates=make(map[int]*Certificate)
			eve.Mapping[fileName]=&eve.Certificates
		case "typeIDs":
			eve.TypeIDs=make(map[int]*TypeID)
			eve.Mapping[fileName]=&eve.TypeIDs
		case "graphicIDs":
			eve.GraphicIDs=make(map[int]*GraphicID)
			eve.Mapping[fileName]=&eve.GraphicIDs
		case "groupIDs":
			eve.GroupIDs=make(map[int]*GroupID)
			eve.Mapping[fileName]=&eve.GroupIDs
		case "iconIDs":
			eve.IconIDs=make(map[int]*IconID)
			eve.Mapping[fileName]=&eve.IconIDs
		case "skins":
			eve.Skins=make(map[int]*Skin)
			eve.Mapping[fileName]=&eve.Skins
		}
	}
	if err = loadStaticData(); err != nil {
		log.Fatalf("Can't load static Data: %s", err)
	}
}
