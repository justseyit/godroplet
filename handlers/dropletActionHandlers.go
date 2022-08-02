package handlers

import (
	"context"
	"godroplet/constants"

	"github.com/digitalocean/godo"
)

func GetActionsByDropletID(dropletID int, options godo.ListOptions) ([]godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	actions, response, err := client.Droplets.Actions(ctx, dropletID, &options)

	return actions, *response, err
}

// Enable Backups
func EnableBackupsByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.EnableBackups(ctx, 3164450)

	return *action, *response, err
}

//Disable Backups
func DisableBackupsByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.DisableBackups(ctx, dropletID)

	return *action, *response, err
}

//Reboot a Droplet
func RebootByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.Reboot(ctx, dropletID)

	return *action, *response, err
}

//Power Cycle a Droplet
func PowerCycleByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.PowerCycle(ctx, dropletID)

	return *action, *response, err
}

//Shutdown a Droplet
func ShutdownByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.Shutdown(ctx, dropletID)

	return *action, *response, err
}

//Power Off a Droplet
func PowerOffByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.PowerOff(ctx, dropletID)

	return *action, *response, err
}

//Power On a Droplet
func PowerOnByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.PowerOn(ctx, dropletID)

	return *action, *response, err
}

//Restore a Droplet
func RestoreByDropletID(dropletID int, imageID int) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.Restore(ctx, dropletID, imageID)

	return *action, *response, err
}

//Password Reset a Droplet
func PasswordResetByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.PasswordReset(ctx, dropletID)

	return *action, *response, err
}

//Resize a Droplet
func ResizeByDropletID(dropletID int, size string, val bool) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.Resize(ctx, dropletID, size, val)

	return *action, *response, err
}

//Rebuild a Droplet
func RebuildByDropletID(dropletID int, imageSlug string) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.RebuildByImageSlug(ctx, dropletID, imageSlug)

	return *action, *response, err
}

//Rename a Droplet
func RenameByDropletID(dropletID int, name string) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.Rename(ctx, dropletID, name)

	return *action, *response, err
}

//Change a Droplet's Kernel
func ChangeKernelByDropletID(dropletID int, kernel int) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.ChangeKernel(ctx, dropletID, kernel)

	return *action, *response, err
}

//Enable IPv6 for a Droplet
func EnableIPv6ByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.EnableIPv6(ctx, dropletID)

	return *action, *response, err
}

//Enable Private Networking for a Droplet
func EnablePrivateNetworkingByDropletID(dropletID int) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.EnablePrivateNetworking(ctx, dropletID)

	return *action, *response, err
}

//Snapshot a Droplet
func SnapshotByDropletID(dropletID int, name string) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.Snapshot(ctx, dropletID, name)

	return *action, *response, err
}

//Retrieve a Droplet Action
func RetrieveByDropletID(dropletID int, actionID int) (godo.Action, godo.Response, error) {

	client := godo.NewFromToken(constants.DIGITALOCEAN_TOKEN)
	ctx := context.TODO()

	action, response, err := client.DropletActions.Get(ctx, dropletID, actionID)

	return *action, *response, err
}
