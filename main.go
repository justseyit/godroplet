package main

import (
	. "godroplet/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter().StrictSlash(true)
	mux.HandleFunc("/newdroplet", PostDropletHandler).Methods("POST")
	mux.HandleFunc("/droplets", GetDropletsHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}", GetDropletHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}", DeleteDropletHandler).Methods("DELETE")
	mux.HandleFunc("/droplets/{dropletID}/actions", GetDropletActionsHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/backups", GetBackupsHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/actions/{action}", PostDropletActionHandler).Methods("POST")
	mux.HandleFunc("/", HomeHandler)

	http.ListenAndServe(":9000", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
