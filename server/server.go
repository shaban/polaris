package server

import (
	"encoding/json"
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
func handleMarketGroups(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bc := db.MarketGroups.Breadcrumb(1674)
	if bc == nil{
		http.Error(w,"Not Found",404)
		return
	}
	w.Header().Set("content-type", "application/json")
	enc := json.NewEncoder(w)
	err := enc.Encode(bc)
	if err != nil{
		http.Error(w,err.Error(),500)
	}
}
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

	gz := gzip.NewWriter(w)
	defer gz.Close()

	if err = db.EncodeByTableAndKey(gz, table, id); err != nil {
		w.Header().Del("content-encoding")
		w.Header().Set("content-type", "text/plain")
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
	router.GET("/", handleMarketGroups)
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
