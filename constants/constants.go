package constants

import (
	"context"

	"github.com/digitalocean/godo"
)

const DIGITALOCEAN_TOKEN string = "dop_v1_6b1de3d9b65a1ac04dc321adcbf3fc3fe3d610e9d1bf73c2540650db9f87e647"

var CLIENT godo.Client
var CTX context.Context

func Initialize() {
	CLIENT = *godo.NewFromToken(DIGITALOCEAN_TOKEN)
	CTX = context.TODO()
}
