package services

import (
	"godroplet/constants"
	"godroplet/utils"
	"net/http"
	"strconv"

	"github.com/digitalocean/godo"
)

// Get a Droplet's Action
func GetActionsByDropletID(dropletID int, options godo.ListOptions) ([]godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	actions, response, err := client.Droplets.Actions(ctx, dropletID, &options)

	return actions, *response, err
}

// Enable Backups
func EnableBackupsByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.EnableBackups(ctx, 3164450)

	return *action, *response, err
}

// Disable Backups
func DisableBackupsByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.DisableBackups(ctx, dropletID)

	return *action, *response, err
}

// Reboot a Droplet
func RebootByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.Reboot(ctx, dropletID)

	return *action, *response, err
}

// Power Cycle a Droplet
func PowerCycleByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.PowerCycle(ctx, dropletID)

	return *action, *response, err
}

// Shutdown a Droplet
func ShutdownByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.Shutdown(ctx, dropletID)

	return *action, *response, err
}

// Power Off a Droplet
func PowerOffByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.PowerOff(ctx, dropletID)

	return *action, *response, err
}

// Power On a Droplet
func PowerOnByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.PowerOn(ctx, dropletID)

	return *action, *response, err
}

// Restore a Droplet
func RestoreByDropletID(dropletID int, imageID int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.Restore(ctx, dropletID, imageID)

	return *action, *response, err
}

// Password Reset a Droplet
func PasswordResetByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.PasswordReset(ctx, dropletID)

	return *action, *response, err
}

// Resize a Droplet
func ResizeByDropletID(dropletID int, size string, val bool) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.Resize(ctx, dropletID, size, val)

	return *action, *response, err
}

// Rebuild a Droplet
func RebuildByDropletID(dropletID int, imageSlug string) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.RebuildByImageSlug(ctx, dropletID, imageSlug)

	return *action, *response, err
}

// Rename a Droplet
func RenameByDropletID(dropletID int, name string) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.Rename(ctx, dropletID, name)

	return *action, *response, err
}

// Change a Droplet's Kernel
func ChangeKernelByDropletID(dropletID int, kernel int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.ChangeKernel(ctx, dropletID, kernel)

	return *action, *response, err
}

// Enable IPv6 for a Droplet
func EnableIPv6ByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.EnableIPv6(ctx, dropletID)

	return *action, *response, err
}

//Disable IPv6 for a Droplet (Have not an endpoint.)
/*func DisableIPv6ByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.DisableIPv6(ctx, dropletID)

	return *action, *response, err
}*/

// Enable Private Networking for a Droplet
func EnablePrivateNetworkingByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.EnablePrivateNetworking(ctx, dropletID)

	return *action, *response, err
}

//Disable Private Networking for a Droplet (Have not an endpoint.)
/*func DisablePrivateNetworkingByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.DisablePrivateNetworking(ctx, dropletID)

	return *action, *response, err
}*/

// Snapshot a Droplet
func SnapshotByDropletID(dropletID int, name string) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.Snapshot(ctx, dropletID, name)

	return *action, *response, err
}

// Retrieve a Droplet Action
func RetrieveByDropletID(dropletID int, actionID int) (godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.Get(ctx, dropletID, actionID)

	return *action, *response, err
}

// Power Cycle by Tag
func PowerCycleByTag(tag string) ([]godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.PowerCycleByTag(ctx, tag)

	return action, *response, err
}

// Power Off by Tag
func PowerOffByTag(tag string) ([]godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.PowerOffByTag(ctx, tag)

	return action, *response, err
}

// Power On by Tag
func PowerOnByTag(tag string) ([]godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.PowerOnByTag(ctx, tag)

	return action, *response, err
}

// Shutdown by Tag
func ShutdownByTag(tag string) ([]godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.ShutdownByTag(ctx, tag)

	return action, *response, err
}

// Enable IPv6 by Tag
func EnableIPv6ByTag(tag string) ([]godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.EnableIPv6ByTag(ctx, tag)

	return action, *response, err
}

// Enable Backups by Tag
func EnableBackupsByTag(tag string) ([]godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.EnableBackupsByTag(ctx, tag)

	return action, *response, err
}

// Disable Backups by Tag
func DisableBackupsByTag(tag string) ([]godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.DisableBackupsByTag(ctx, tag)
	return action, *response, err
}

// Snapshot by Tag
func SnapshotByTag(tag string, name string) ([]godo.Action, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	action, response, err := client.DropletActions.SnapshotByTag(ctx, tag, name)
	return action, *response, err
}

// Retrieve a Droplet Action by Droplet ID and Action ID
func RetrieveByDropletIDAndActionID(dropletID int, actionID int) (http.Response, error) {

	//client := constants.CLIENT
	ctx := constants.CTX

	req, err := http.Get(constants.URL + "/droplets/" + strconv.Itoa(dropletID) + "/actions/" + strconv.Itoa(actionID))

	req.Header.Set("Authorization", "Bearer "+constants.DIGITALOCEAN_TOKEN)
	req.Header.Set("Content-Type", "application/json")
	utils.CheckError(err)

	response, err := godo.DoRequest(ctx, req.Request)

	return *response, err
}
