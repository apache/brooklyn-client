package version

import (
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func Version(network *net.Network) string {
	url := "/v1/version"
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}
