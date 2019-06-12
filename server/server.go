package server

import (
	"compress/gzip"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/shaban/polaris/db"
	"github.com/spf13/viper"
	"golang.org/x/crypto/acme/autocert"
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
	w.Header().Set("content-encoding", "gzip")
	w.Header().Set("server", "Kengal 2.0")

	gzip := gzip.NewWriter(w)
	defer gzip.Close()

	if err = db.EncodeByTableAndKey(gzip, table, id); err != nil {
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
	err := http.Serve(autocert.NewListener(address), router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func setDebugOn() {
	debug = true
}
func setDebugOff() {
	debug = false
}
