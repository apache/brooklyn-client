package commands

import (
	"github.com/apache/brooklyn-client/net"
)

type CatalogEntities struct {
	network *net.Network
}

func NewCatalogEntities(network *net.Network) (cmd *CatalogEntities) {
	cmd = new(CatalogEntities)
	cmd.network = network
	return
}
