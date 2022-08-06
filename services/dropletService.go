package services

import (
	"godroplet/constants"
	"godroplet/utils"

	"github.com/digitalocean/godo"
)

/*createRequest := &godo.DropletCreateRequest{
Name:   "example.com",
Region: "nyc3",
Size:   "s-1vcpu-1gb",
Image: godo.DropletCreateImage{
	Slug: "ubuntu-16-04-x64",
},
SSHKeys: []godo.DropletCreateSSHKey{
	{ID: 107149},
},
IPv6: true,
Tags: []string{"web"},}*/

func CreateNewDroplet(createRequest godo.DropletCreateRequest) (godo.Droplet, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	droplet, response, err := client.Droplets.Create(ctx, &createRequest)

	utils.CheckError(err)
	return *droplet, *response, err
}

func CreateNewDroplets(createRequest godo.DropletMultiCreateRequest) ([]godo.Droplet, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	droplets, response, err := client.Droplets.CreateMultiple(ctx, &createRequest)

	utils.CheckError(err)
	return droplets, *response, err
}

func ListAllDroplets(options *godo.ListOptions) ([]godo.Droplet, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	droplets, response, err := client.Droplets.List(ctx, options)
	return droplets, *response, err

}

func DeleteDropletByTag(tag string) (godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	response, err := client.Droplets.DeleteByTag(ctx, tag)

	if err != nil {
		return *response, err
	} else {
		return *response, err
	}
}

func GetDropletByID(id int) (godo.Droplet, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	droplet, response, err := client.Droplets.Get(ctx, id)

	return *droplet, *response, err
}

func DeleteDropletByID(id int) (godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	response, err := client.Droplets.Delete(ctx, id)
	if err != nil {
		return *response, err
	} else {
		return *response, err
	}
}

func DeleteDropletsByTag(tag string) (godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	response, err := client.Droplets.DeleteByTag(ctx, tag)
	if err != nil {
		return *response, err
	} else {
		return *response, err
	}
}

func GetBackupsByDropletID(id int, options *godo.ListOptions) ([]godo.Image, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	backups, response, err := client.Droplets.Backups(ctx, id, options)

	return backups, *response, err
}

func GetsSnapshotsByDropletID(id int, options *godo.ListOptions) ([]godo.Image, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	snapshots, response, err := client.Droplets.Snapshots(ctx, id, options)

	return snapshots, *response, err
}

func GetAvailableKernelsByDropletID(id int, options *godo.ListOptions) ([]godo.Kernel, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	kernels, response, err := client.Droplets.Kernels(ctx, id, options)

	return kernels, *response, err
}

// /****************************************************************************************************///
func GetNeighborsByDropletID(id int, options *godo.ListOptions) ([]godo.Droplet, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	neighbors, response, err := client.Droplets.Neighbors(ctx, id)

	return neighbors, *response, err
}

func GetAppliedFirewallsByDropletID(id int, options *godo.ListOptions) ([]godo.Firewall, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	firewalls, response, err := client.Firewalls.ListByDroplet(ctx, id, options)

	return firewalls, *response, err
}

/*
func GetAssociatedResourcesByDropletID(id int, options *godo.ListOptions) (godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX


	req, err := http.NewRequest("GET", "https://api.digitalocean.com/v2/droplets/"+string(id)+"/resources", nil)

	response, err := godo.DoRequest(ctx, req)

	return  *response, err
}*/
