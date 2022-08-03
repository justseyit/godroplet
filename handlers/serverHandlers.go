package handlers

import (
	"godroplet/utils"
	"net/http"
	"strconv"

	"github.com/digitalocean/godo"
	"github.com/gorilla/mux"
)

//HTTP Handler for getting a droplet by droplet id
func GetDropletHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)
	droplet, _, err := GetDropletByID(dropletIDInt)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(droplet.String()))
}

//HTTP Handler for getting all droplets
func GetDropletsHandler(w http.ResponseWriter, r *http.Request) {
	options := godo.ListOptions{}
	droplets, _, err := ListAllDroplets(&options)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for _, droplet := range droplets {
		w.Write([]byte(droplet.String()))
	}
}

//HTTP Handler for creating a droplet
func PostDropletHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	//Get the form values
	name := vars["name"]
	size := vars["size"]
	image := vars["image"]
	region := vars["region"]
	sshKey, _ := strconv.Atoi(vars["ssh_key"])
	backups, _ := strconv.ParseBool(vars["backups"])
	ipv6, _ := strconv.ParseBool(vars["ipv6"])
	privateNetworking, _ := strconv.ParseBool(vars["private_networking"])
	userData := vars["user_data"]

	//Create the droplet
	createRequest := godo.DropletCreateRequest{
		Name:   name,
		Region: region,
		Size:   size,
		Image: godo.DropletCreateImage{
			Slug: image,
		},
		SSHKeys: []godo.DropletCreateSSHKey{
			{ID: sshKey},
		},
		Backups:           backups,
		IPv6:              ipv6,
		PrivateNetworking: privateNetworking,
		UserData:          userData,
	}

	droplet, _, err := CreateNewDroplet(createRequest)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(droplet.String()))
}

func GetDropletActionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)
	options := godo.ListOptions{}
	actions, _, err := GetActionsByDropletID(dropletIDInt, options)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for _, action := range actions {
		w.Write([]byte(action.String()))
	}
}

func DeleteDropletHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)
	_, err := DeleteDropletByID(dropletIDInt)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetBackupsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)
	options := godo.ListOptions{}
	backups, _, err := GetBackupsByDropletID(dropletIDInt, &options)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	for _, backup := range backups {
		w.Write([]byte(backup.String()))
	}
}
