package commands

import (
	"github.com/apache/brooklyn-client/net"
)

type CatalogApplication struct {
	network *net.Network
}

func NewCatalogApplication(network *net.Network) (cmd *CatalogApplication) {
	cmd = new(CatalogApplication)
	cmd.network = network
	return
}
