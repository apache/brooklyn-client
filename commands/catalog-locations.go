package commands

import (
	"github.com/brooklyncentral/brooklyn-cli/net"
)

type CatalogLocations struct {
	network *net.Network
}

func NewCatalogLocations(network *net.Network) (cmd *CatalogLocations) {
	cmd = new(CatalogLocations)
	cmd.network = network
	return
}