package commands

import (
	"github.com/apache/brooklyn-client/net"
)

type DeleteCatalogApplication struct {
	network *net.Network
}

func NewDeleteCatalogApplication(network *net.Network) (cmd *DeleteCatalogApplication) {
	cmd = new(DeleteCatalogApplication)
	cmd.network = network
	return
}
