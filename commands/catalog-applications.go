package commands

import (
	"github.com/brooklyncentral/brooklyn-cli/net"
)

type CatalogApplication struct {
	network *net.Network
}

func NewCatalogApplication(network *net.Network) (cmd *CatalogApplication) {
	cmd = new(CatalogApplication)
	cmd.network = network
	return
}
