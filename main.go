package main

import (
	"godroplet/constants"
	. "godroplet/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	constants.Initialize()
	mux := mux.NewRouter().StrictSlash(true)
	mux.HandleFunc("/droplets", GetDropletsHandler).Methods("GET")
	mux.HandleFunc("/droplets/single", PostDropletHandler).Methods("POST")
	mux.HandleFunc("/droplets/multiple", PostDropletsHandler).Methods("POST")
	mux.HandleFunc("/droplets/{dropletID}", GetDropletHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}", DeleteDropletHandler).Methods("DELETE")
	mux.HandleFunc("/droplets/{tag}", DeleteDropletByTagHandler).Methods("DELETE")
	mux.HandleFunc("/droplets/{dropletID}/actions", GetDropletActionsHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/backups", GetBackupsHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/actions/{action}", PostDropletActionHandler).Methods("POST")
	mux.HandleFunc("/droplets/{tag}/actions/{action}", PostDropletActionByTagHandler).Methods("POST")
	mux.HandleFunc("/droplets/{dropletID}/snapshots", GetDropletSnapshotsHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/kernels", GetDropletKernelsHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/resources", GetDropletResourcesHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/neighbors", GetDropletNeighborsHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/firewalls", GetFirewallsHandler).Methods("GET")
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/introduction", HomeHandler)

	http.ListenAndServe(":9000", mux)
}

// HTTP Handler for the introduction page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
