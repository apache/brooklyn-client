package commands

import (
	"github.com/brooklyncentral/brooklyn-cli/net"
)

type CatalogEntities struct {
	network *net.Network
}

func NewCatalogEntities(network *net.Network) (cmd *CatalogEntities) {
	cmd = new(CatalogEntities)
	cmd.network = network
	return
}
