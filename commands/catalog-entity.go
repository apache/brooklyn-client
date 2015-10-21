package commands

import (
	"github.com/brooklyncentral/brooklyn-cli/net"
)

type CatalogEntity struct {
	network *net.Network
}

func NewCatalogEntity(network *net.Network) (cmd *CatalogEntity) {
	cmd = new(CatalogEntity)
	cmd.network = network
	return
}