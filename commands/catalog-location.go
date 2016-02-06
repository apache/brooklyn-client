package commands

import (
	"github.com/brooklyncentral/brooklyn-cli/net"
)

type CatalogLocation struct {
	network *net.Network
}

func NewCatalogLocation(network *net.Network) (cmd *CatalogLocation) {
	cmd = new(CatalogLocation)
	cmd.network = network
	return
}
