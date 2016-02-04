package commands

import (
	"github.com/brooklyncentral/brooklyn-cli/net"
)

type DeleteCatalogApplication struct {
	network *net.Network
}

func NewDeleteCatalogApplication(network *net.Network) (cmd *DeleteCatalogApplication) {
	cmd = new(DeleteCatalogApplication)
	cmd.network = network
	return
}