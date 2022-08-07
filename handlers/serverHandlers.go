package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	. "godroplet/services"
	"godroplet/utils"
	"net/http"
	"strconv"

	"github.com/digitalocean/godo"
	"github.com/gorilla/mux"
)

// HTTP Handler for getting a droplet by droplet id
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

// HTTP Handler for getting all droplets
func GetDropletsHandler(w http.ResponseWriter, r *http.Request) {
	options := godo.ListOptions{}
	droplets, response, err := ListAllDroplets(&options)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	data, _ := json.Marshal(droplets)
	w.Write([]byte(data))
}

// HTTP Handler for creating a droplet
func PostDropletHandler(w http.ResponseWriter, r *http.Request) {

	var createRequest godo.DropletCreateRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createRequest)
	utils.CheckErrorAsResponse(err, w)

	fmt.Println(createRequest)

	droplet, response, err := CreateNewDroplet(createRequest)
	utils.CheckErrorAsResponse(err, w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	data, _ := json.Marshal(droplet)
	w.Write([]byte(data))
}

// HTTP Handler for getting a droplet's actions by droplet id
func GetDropletActionsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)
	options := godo.ListOptions{}
	actions, response, err := GetActionsByDropletID(dropletIDInt, options)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	data, _ := json.Marshal(actions)
	w.Write([]byte(data))
}

// HTTP Handler for create multiple droplets
func PostDropletsHandler(w http.ResponseWriter, r *http.Request) {
	var createRequest godo.DropletMultiCreateRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&createRequest)
	utils.CheckErrorAsResponse(err, w)

	fmt.Println(createRequest)

	droplets, response, err := CreateNewDroplets(createRequest)
	utils.CheckErrorAsResponse(err, w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	data, _ := json.Marshal(droplets)
	w.Write([]byte(data))
}

// HTTP Handler for deleting a droplet by droplet id
func DeleteDropletHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)
	response, err := DeleteDropletByID(dropletIDInt)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
}

// HTTP Handler for deleting droplets by tag
func DeleteDropletByTagHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tag := vars["tag"]
	response, err := DeleteDropletsByTag(tag)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
}

// HTTP Handler for getting a droplet's backups by droplet id
func GetBackupsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)
	options := godo.ListOptions{}
	backups, response, err := GetBackupsByDropletID(dropletIDInt, &options)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	data, _ := json.Marshal(backups)
	w.Write([]byte(data))
}

// HTTP Handler for getting a droplet's kernels by droplet id
func GetDropletKernelsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)
	options := godo.ListOptions{}
	kernels, response, err := GetAvailableKernelsByDropletID(dropletIDInt, &options)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	data, _ := json.Marshal(kernels)
	w.Write([]byte(data))
}

// HTTP Handler for getting a droplet's snapshots by droplet id
func GetDropletSnapshotsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)
	options := godo.ListOptions{}
	snapshots, response, err := GetsSnapshotsByDropletID(dropletIDInt, &options)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	data, _ := json.Marshal(snapshots)
	w.Write([]byte(data))
}

// HTTP Handler for getting a droplet's neighbors by droplet id
func GetDropletNeighborsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)
	options := godo.ListOptions{}
	neighbors, response, err := GetNeighborsByDropletID(dropletIDInt, &options)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	data, _ := json.Marshal(neighbors)
	w.Write([]byte(data))
}

// HTTP Handler for getting associated resources by droplet id
func GetDropletResourcesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)
	options := godo.ListOptions{}
	response, err := GetAssociatedResourcesByDropletID(dropletIDInt, &options)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	data, _ := json.Marshal(response.Body)
	w.Write([]byte(data))
}

// HTTP Handler for getting firewalls applied to a droplet by droplet id
func GetFirewallsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)
	options := godo.ListOptions{}
	firewalls, response, err := GetAppliedFirewallsByDropletID(dropletIDInt, &options)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	data, _ := json.Marshal(firewalls)
	w.Write([]byte(data))
}

// HTTP Handler for deleting a droplet with associated resources selectively
func DeleteDropletWithResourcesSelectivelyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)

	decoder := json.NewDecoder(r.Body)

	var data string

	decoder.Decode(&data)

	response, err := SelectivelyDeleteDropletWithAssociatedResourcesByID(dropletIDInt, data)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
}

// HTTP Handler for deleting a droplet with associated resources dangerous
func DeleteDropletWithResourcesDangerousHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)

	alertStr := r.Header.Get("alert")
	alert, err := strconv.ParseBool(alertStr)
	utils.CheckErrorAsResponse(err, w)

	response, err := DestroyDropletWithAssociatedResourcesByID(dropletIDInt, alert)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
}

// HTTP Handler for getting a droplet's neighbor id list
func GetDropletNeighborIDsHandler(w http.ResponseWriter, r *http.Request) {

	response, err := GetAllDropletNeighborsIDs()
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	data, _ := json.Marshal(response.Body)

	w.Write([]byte(data))
}

// HTTP Handler for Check Status of a Droplet Destroy with Associated Resources Request
func CheckDropletDestroyStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)

	response, err := CheckStatusOfDropletDestroyRequest(dropletIDInt)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	data, _ := json.Marshal(response.Body)

	w.Write([]byte(data))
}

// HTTP Handler for Retry a Droplet Destroy with Associated Resources Request
func RetryDropletDestroyRequestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)

	response, err := RetryDropletDestroyRequest(dropletIDInt)
	utils.CheckError(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	data, _ := json.Marshal(response.Body)

	w.Write([]byte(data))
}

// HTTP Handler for getting a droplet's action by droplet id and action id
func GetDropletActionByActionIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	dropletIDInt, _ := strconv.Atoi(dropletID)
	actionID := vars["actionID"]
	actionIDInt, _ := strconv.Atoi(actionID)
	response, err := RetrieveByDropletIDAndActionID(dropletIDInt, actionIDInt)
	utils.CheckErrorAsResponse(err, w)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	data, _ := json.Marshal(response.Body)
	w.Write([]byte(data))
}

// HTTP Handler for initiate an action by droplet id
func PostDropletActionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dropletID := vars["dropletID"]
	actionStr := vars["action"]
	fmt.Println(vars)
	dropletIDInt, _ := strconv.Atoi(dropletID)

	switch actionStr {
	case "power_cycle":
		action, response, err := PowerCycleByDropletID(dropletIDInt)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "power_off":
		action, response, err := PowerOffByDropletID(dropletIDInt)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "power_on":
		action, response, err := PowerOnByDropletID(dropletIDInt)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "reboot":
		action, response, err := RebootByDropletID(dropletIDInt)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "shutdown":
		action, response, err := ShutdownByDropletID(dropletIDInt)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "snapshot":
		var name string
		err := r.ParseForm()
		name = r.FormValue("name")
		utils.CheckErrorAsResponse(err, w)

		action, response, err := SnapshotByDropletID(dropletIDInt, name)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "restore":
		var imageID int
		err := r.ParseForm()
		imageID, _ = strconv.Atoi(r.FormValue("image_id"))
		utils.CheckErrorAsResponse(err, w)

		action, response, err := RestoreByDropletID(dropletIDInt, imageID)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "password_reset":
		action, response, err := PasswordResetByDropletID(dropletIDInt)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "rebuild":
		var imageSlug string
		err := r.ParseForm()
		imageSlug = r.FormValue("image_slug")
		utils.CheckErrorAsResponse(err, w)

		action, response, err := RebuildByDropletID(dropletIDInt, imageSlug)
		utils.CheckErrorAsResponse(err, w)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "resize":
		var size string
		var val bool
		err := r.ParseForm()
		size = r.FormValue("size")
		val, _ = strconv.ParseBool(r.FormValue("val"))
		utils.CheckErrorAsResponse(err, w)

		action, response, err := ResizeByDropletID(dropletIDInt, size, val)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "rename":
		var name string
		err := r.ParseForm()
		name = r.FormValue("name")
		utils.CheckErrorAsResponse(err, w)
		//Get name value from form

		action, response, err := RenameByDropletID(dropletIDInt, name)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "change_kernel":
		var kernelID int
		err := r.ParseForm()
		kernelID, _ = strconv.Atoi(r.FormValue("kernel_id"))
		utils.CheckErrorAsResponse(err, w)

		action, response, err := ChangeKernelByDropletID(dropletIDInt, kernelID)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "enable_ipv6":
		action, response, err := EnableIPv6ByDropletID(dropletIDInt)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	/*case "disable_ipv6":
	action, response, err := DisableIPv6ByDropletID(dropletIDInt)
	break*/
	case "enable_private_networking":
		action, response, err := EnablePrivateNetworkingByDropletID(dropletIDInt)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	/*case "disable_private_networking":
	_, err := DisablePrivateNetworkingDropletByID(dropletIDInt)
	break*/
	case "enable_backups":
		action, response, err := EnableBackupsByDropletID(dropletIDInt)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "disable_backups":
		action, response, err := DisableBackupsByDropletID(dropletIDInt)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	default:
		err := errors.New("Invalid action")
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		break
	}
}

// HTTP Handler for initiate an action by tags
func PostDropletActionByTagHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tag := vars["tag"]
	actionStr := vars["action"]

	switch actionStr {
	case "snapshot":
		var name string
		err := r.ParseForm()
		name = r.FormValue("name")
		utils.CheckErrorAsResponse(err, w)
		action, response, err := SnapshotByTag(tag, name)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "shutdown":
		action, response, err := ShutdownByTag(tag)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "power_off":
		action, response, err := PowerOffByTag(tag)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "power_on":
		action, response, err := PowerOnByTag(tag)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "power_cycle":
		action, response, err := PowerCycleByTag(tag)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "enable_ipv6":
		action, response, err := EnableIPv6ByTag(tag)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "enable_backups":
		action, response, err := EnableBackupsByTag(tag)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	case "disable_backups":
		action, response, err := DisableBackupsByTag(tag)
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		data, _ := json.Marshal(action)
		w.Write([]byte(data))
		break
	default:
		err := errors.New("Invalid action")
		utils.CheckErrorAsResponse(err, w)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		break
	}
}
