package commands

import (
	"github.com/apache/brooklyn-client/net"
)

type CatalogEntity struct {
	network *net.Network
}

func NewCatalogEntity(network *net.Network) (cmd *CatalogEntity) {
	cmd = new(CatalogEntity)
	cmd.network = network
	return
}
