package commands

import (
	"github.com/brooklyncentral/brooklyn-cli/net"
)

type DeleteCatalogEntity struct {
	network *net.Network
}

func NewDeleteCatalogEntity(network *net.Network) (cmd *DeleteCatalogEntity) {
	cmd = new(DeleteCatalogEntity)
	cmd.network = network
	return
}
