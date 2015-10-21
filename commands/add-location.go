package commands

import (
	"github.com/brooklyncentral/brooklyn-cli/net"
)

type AddLocation struct {
	network *net.Network
}

func NewAddLocation(network *net.Network) (cmd *AddLocation) {
	cmd = new(AddLocation)
	cmd.network = network
	return
}