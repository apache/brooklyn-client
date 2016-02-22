package commands

import (
	"github.com/apache/brooklyn-client/net"
)

type AddLocation struct {
	network *net.Network
}

func NewAddLocation(network *net.Network) (cmd *AddLocation) {
	cmd = new(AddLocation)
	cmd.network = network
	return
}
