package commands

import (
	"github.com/apache/brooklyn-client/net"
)

type ResetCatalog struct {
	network *net.Network
}

func NewResetCatalog(network *net.Network) (cmd *ResetCatalog) {
	cmd = new(ResetCatalog)
	cmd.network = network
	return
}
