package main

import (
	_ "github.com/lib/pq"
	"github.com/shaban/polaris/db"
	"github.com/shaban/polaris/server"
	"github.com/spf13/viper"
	"log"
)

//go:generate sh db/generate.sh
var(
	conf *viper.Viper
)

func main() {
	var(
		err error
	)
	conf := viper.New()
	conf.SetConfigName("config")
	conf.AddConfigPath("$HOME/code/src/github.com/shaban/polaris")
	if err = conf.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	db.Open(conf)
	defer db.Close()
	server.Start(conf)
}
