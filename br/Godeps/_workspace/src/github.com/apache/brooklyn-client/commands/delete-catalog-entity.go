package commands

import (
	"github.com/apache/brooklyn-client/net"
)

type DeleteCatalogEntity struct {
	network *net.Network
}

func NewDeleteCatalogEntity(network *net.Network) (cmd *DeleteCatalogEntity) {
	cmd = new(DeleteCatalogEntity)
	cmd.network = network
	return
}
