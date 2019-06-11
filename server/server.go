package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/shaban/polaris/db"
	"github.com/spf13/viper"
)

var (
	port    string
	address string
	debug   bool
)

func handleAPI(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		table string
		id    int
		err   error
		//out   []byte
	)
	table = ps.ByName("table")
	idString := ps.ByName("id")

	if id, err = strconv.Atoi(idString); err != nil {
		http.Error(w, "Not Found!", 404)
	}
	w.Header().Set("content-type", "application/json")

	if err = db.EncodeByTableAndKey(w, table, id); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

//Start gets the relevant info from the configuration file
//registers the routes
//and starts the server
func Start(conf *viper.Viper) {
	address = conf.GetString("host.address")
	port = conf.GetString("host.port")
	router := httprouter.New()
	router.GET("/api/:table/:id", handleAPI)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)) //136.243.94.189
	//log.Fatal(http.ListenAndServe("136.243.94.189:8081", router))//136.243.94.189
}

func setDebugOn() {
	debug = true
}
func setDebugOff() {
	debug = false
}
