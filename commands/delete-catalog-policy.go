package commands

import (
	"github.com/apache/brooklyn-client/net"
)

type DeleteCatalogPolicy struct {
	network *net.Network
}

func NewDeleteCatalogPolicy(network *net.Network) (cmd *DeleteCatalogPolicy) {
	cmd = new(DeleteCatalogPolicy)
	cmd.network = network
	return
}
