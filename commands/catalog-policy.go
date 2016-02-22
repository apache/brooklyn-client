package commands

import (
	"github.com/apache/brooklyn-client/net"
)

type CatalogPolicy struct {
	network *net.Network
}

func NewCatalogPolicy(network *net.Network) (cmd *CatalogPolicy) {
	cmd = new(CatalogPolicy)
	cmd.network = network
	return
}
