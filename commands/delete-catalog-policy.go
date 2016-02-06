package commands

import (
	"github.com/brooklyncentral/brooklyn-cli/net"
)

type DeleteCatalogPolicy struct {
	network *net.Network
}

func NewDeleteCatalogPolicy(network *net.Network) (cmd *DeleteCatalogPolicy) {
	cmd = new(DeleteCatalogPolicy)
	cmd.network = network
	return
}
