# godroplet

## [EN]

This tool has allows you to use Digital Ocean Droplet and Droplet Actions API.

## HTTP GET Methods

* `/droplets`: Get all droplets.
* `/droplets/{dropletID}`: Get droplet by id. Replace `{dropletID}` with droplet ID.
* `/droplets/{dropletID}/actions`: Get all actions of a droplet by droplet ID. Replace `{dropletID}` with droplet ID.
* `/droplets/{dropletID}/backups`: Get all backups of a droplet by droplet ID. Replace `{dropletID}` with droplet ID.
* `/droplets/{dropletID}/snapshots`: Get all snapshots of a droplet by droplet ID. Replace `{dropletID}` with droplet ID.
* `/droplets/{dropletID}/kernels`: Get all kernels of a droplet by droplet ID. Replace `{dropletID}` with droplet ID.
* `/droplets/{dropletID}/resources`: Get all resources of a droplet by droplet ID. Replace `{dropletID}` with droplet ID.
* `/droplets/{dropletID}/neighbors`: Get all neighbors of a droplet by droplet ID. Replace `{dropletID}` with droplet ID.
* `/droplets/{dropletID}/firewalls`: Get all firewalls of a droplet by droplet ID. Replace `{dropletID}` with droplet ID.
* `/droplets/neighbors`: Get all droplet ID lists that co-located on the same physical hardware.
* `/droplets/{dropletID}/resources/status`: Check status of a droplet destroy with associated resources request by droplet ID. Replace `{dropletID}` with droplet ID.
* `/droplets/{dropletID}/actions/{actionID}`: Get an action on a droplet by droplet ID and action ID. Replace `{dropletID}` with droplet ID and `{actionID}` with action ID.

## HTTP POST Methods

* `/droplets/single`: Create a droplet. In the request body, you need to specify droplet properties as JSON. You can find the properties in [this part of Digital Ocean Documentation](https://docs.digitalocean.com/reference/api/api-reference/#operation/droplets_create).
*`/droplets/multiple`: Create multiple droplets. In the request body, you need to specify droplet properties as JSON. You can find the properties in [this part of Digital Ocean Documentation](https://docs.digitalocean.com/reference/api/api-reference/#operation/droplets_create).
* `/droplets/{dropletID}/actions/{action}`: Perform an action on a droplet by droplet ID and action. Replace `{dropletID}` with droplet ID and `{action}` with action. You can find the actions in [this part of Digital Ocean Documentation](https://docs.digitalocean.com/reference/api/api-reference/#operation/dropletActions_post).
* `/droplets/tags/{tag}/actions/{action}`: Perform an action on a droplet by tag and action. Replace `{tag}` with tag and `{action}` with action. You can find the actions in [this part of Digital Ocean Documentation](https://docs.digitalocean.com/reference/api/api-reference/#operation/dropletActions_post_byTag).
* `/droplets/{dropletID}/resources/retry`: Retry a "Droplet Destroy with Associated Resources Request" by droplet ID. Replace `{dropletID}` with droplet ID.

## HTTP DELETE Methods

* `/droplets/{dropletID}`: Delete a droplet by droplet ID. Replace `{dropletID}` with droplet ID.
* `/droplets/tags/{tag}`: Delete a droplet by tag. Replace `{tag}` with tag.
* `/droplets/{dropletID}/resources/selective`: Delete a droplet by droplet ID and delete associated resources selectively. Replace `{dropletID}` with droplet ID. You should specify the resources you want to delete in the request body as JSON. You can find the request body schema in [this part of Digital Ocean Documentation](https://docs.digitalocean.com/reference/api/api-reference/#operation/droplets_destroy_withAssociatedResourcesSelective).
* `/droplets/{dropletID}/resources/dangerous`: Delete a droplet by droplet ID and delete all of associated resources. Replace `{dropletID}` with droplet ID. You have to specify the boolean value as string "true" or "false" in the request header. You can find the request header parameters in [this part of Digital Ocean Documentation](https://docs.digitalocean.com/reference/api/api-reference/#operation/droplets_destroy_withAssociatedResourcesDangerous).
