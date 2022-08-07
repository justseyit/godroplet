package main

import (
	"godroplet/constants"
	. "godroplet/handlers"
	"godroplet/utils"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	constants.Initialize()
	mux := mux.NewRouter().StrictSlash(true)
	mux.HandleFunc("/droplets/neighbors", GetDropletNeighborIDsHandler).Methods("GET")
	mux.HandleFunc("/droplets", GetDropletsHandler).Methods("GET")
	mux.HandleFunc("/droplets/single", PostDropletHandler).Methods("POST")
	mux.HandleFunc("/droplets/multiple", PostDropletsHandler).Methods("POST")
	mux.HandleFunc("/droplets/{dropletID}", GetDropletHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}", DeleteDropletHandler).Methods("DELETE")
	mux.HandleFunc("/droplets/tags/{tag}", DeleteDropletByTagHandler).Methods("DELETE")
	mux.HandleFunc("/droplets/{dropletID}/actions", GetDropletActionsHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/actions/{actionID}", GetDropletActionByActionIDHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/backups", GetBackupsHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/actions/{action}", PostDropletActionHandler).Methods("POST")
	mux.HandleFunc("/droplets/tags/{tag}/actions/{action}", PostDropletActionByTagHandler).Methods("POST")
	mux.HandleFunc("/droplets/{dropletID}/snapshots", GetDropletSnapshotsHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/kernels", GetDropletKernelsHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/resources", GetDropletResourcesHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/neighbors", GetDropletNeighborsHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/firewalls", GetFirewallsHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/resources/selective", DeleteDropletWithResourcesSelectivelyHandler).Methods("DELETE")
	mux.HandleFunc("/droplets/{dropletID}/resources/dangerous", DeleteDropletWithResourcesDangerousHandler).Methods("DELETE")
	mux.HandleFunc("/droplets/{dropletID}/resources/status", CheckDropletDestroyStatusHandler).Methods("GET")
	mux.HandleFunc("/droplets/{dropletID}/resources/retry", RetryDropletDestroyRequestHandler).Methods("POST")
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/introduction", HomeHandler)

	http.ListenAndServe(":9000", mux)
}

// HTTP Handler for the introduction page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	//Write to body ./pages/intro.html

	file, err := os.Open("./pages/intro.html")
	utils.CheckErrorAsResponse(err, w)
	defer file.Close()

	// Write to body
	_, err = io.Copy(w, file)
	utils.CheckErrorAsResponse(err, w)

}
