package usage

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func Application(network *net.Network, application string) string {
	url := fmt.Sprintf("/v1/usage/applications/%s", application)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func Applications(network *net.Network) string {
	url := fmt.Sprintf("/v1/usage/applications")
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func Machine(network *net.Network, machine string) string {
	url := fmt.Sprintf("/v1/usage/machines/%s", machine)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func Machines(network *net.Network) string {
	url := fmt.Sprintf("/v1/usage/machines")
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}