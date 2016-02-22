package commands

import (
	"github.com/apache/brooklyn-client/net"
)

type CatalogLocation struct {
	network *net.Network
}

func NewCatalogLocation(network *net.Network) (cmd *CatalogLocation) {
	cmd = new(CatalogLocation)
	cmd.network = network
	return
}
