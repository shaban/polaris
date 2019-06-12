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

	pg               *sql.DB
	eve              *EveDB
)

//Open reads the configuration opens the postgres database
//sets path variables and boostraps the type mapper interface
//finally it loads all static data on an audit from yaml
//otherwise from database into the mapper map
func Open(conf *viper.Viper) {
	var (
		err error
	)
	connectionString = fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s",
		conf.GetString("database.postgres.user"),
		conf.GetString("database.postgres.db"),
		conf.GetString("database.postgres.password"),
		conf.GetString("database.postgres.host"),
	)

	if err = open(); err != nil {
		log.Fatal(err)
	}

	basePath = conf.GetString("path")
	yamlPath = conf.GetString("database.yaml.path")
	yamlExt = conf.GetString("database.yaml.extension")
	yamlFiles = conf.GetStringSlice("database.yaml.files")

	eve = new(EveDB)
	eve.Mapping = make(map[string]mapper)

	for _, fileName := range yamlFiles{
		switch fileName{
		case "blueprints":
			eve.Mapping[fileName]=blueprints(make(map[int]*Blueprint))
		case "categoryIDs":
			eve.Mapping[fileName]=categoryIDs(make(map[int]*CategoryID))
		case "certificates":
			eve.Mapping[fileName]=certificates(make(map[int]*Certificate))
		case "typeIDs":
			eve.Mapping[fileName]=typeIDs(make(map[int]*TypeID))
		case "graphicIDs":
			eve.Mapping[fileName]=graphicIDs(make(map[int]*GraphicID))
		case "groupIDs":
			eve.Mapping[fileName]=groupIDs(make(map[int]*GroupID))
		case "iconIDs":
			eve.Mapping[fileName]=iconIDs(make(map[int]*IconID))
		case "skins":
			eve.Mapping[fileName]=skins(make(map[int]*Skin))
		case "invMarketGroups":
			eve.Mapping[fileName]=marketGroups(make(map[int]*MarketGroup))
		}
	}
	if err = loadStaticData(); err != nil {
		log.Fatalf("Can't load static Data: %s", err)
	}
}
