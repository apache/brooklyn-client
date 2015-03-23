package version

import(
	"fmt"
	"github.com/robertgmoss/brooklyn-cli/net"
)

func Version(network *net.Network) string{
	url := "/v1/version"
	req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}