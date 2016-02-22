package usage

import (
	"encoding/json"
	"fmt"
	"github.com/apache/brooklyn-client/models"
	"github.com/apache/brooklyn-client/net"
)

func Application(network *net.Network, application string) (string, error) {
	url := fmt.Sprintf("/v1/usage/applications/%s", application)
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Applications(network *net.Network) (string, error) {
	url := fmt.Sprintf("/v1/usage/applications")
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Machine(network *net.Network, machine string) (string, error) {
	url := fmt.Sprintf("/v1/usage/machines/%s", machine)
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Machines(network *net.Network) (string, error) {
	url := fmt.Sprintf("/v1/usage/machines")
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
