package version

import(
	"fmt"
	"github.com/robertgmoss/brooklyn-cli/net"
)

func Version() string{
	url := "http://192.168.50.101:8081/v1/version"
	req := net.NewGetRequest(url)
	body, err := net.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}