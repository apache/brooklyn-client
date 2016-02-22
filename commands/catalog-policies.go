package commands

import (
	"github.com/apache/brooklyn-client/net"
)

type CatalogPolicies struct {
	network *net.Network
}

func NewCatalogPolicies(network *net.Network) (cmd *CatalogPolicies) {
	cmd = new(CatalogPolicies)
	cmd.network = network
	return
}
