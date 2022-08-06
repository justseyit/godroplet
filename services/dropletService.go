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

	/*droplets, _, err := client.Droplets.List(ctx, &godo.ListOptions{})
	if err != nil {
		return droplets, err
	} else {
		return droplets, err
	}*/

	droplets, response, err := client.Droplets.List(ctx, options)

	//snapshot, response, err := client.Droplets.Snapshots(ctx, 3164494, options)

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

func GetBackupsByDropletID(id int, options *godo.ListOptions) ([]godo.Image, godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	backups, response, err := client.Droplets.Backups(ctx, id, options)

	return backups, *response, err
}

func GetsSnapshotsByDropletID(id int, options *godo.ListOptions) (godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	_, response, err := client.Droplets.Snapshots(ctx, id, options)
	if err != nil {
		return *response, err
	} else {
		return *response, err
	}
}

func GetAvailableKernelsByDropletID(id int, options *godo.ListOptions) (godo.Response, error) {

	client := constants.CLIENT
	ctx := constants.CTX

	_, response, err := client.Droplets.Kernels(ctx, id, options)
	if err != nil {
		return *response, err
	} else {
		return *response, err
	}
}