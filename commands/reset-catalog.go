package commands

import (
	"github.com/brooklyncentral/brooklyn-cli/net"
)

type ResetCatalog struct {
	network *net.Network
}

func NewResetCatalog(network *net.Network) (cmd *ResetCatalog) {
	cmd = new(ResetCatalog)
	cmd.network = network
	return
}
