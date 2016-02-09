package server

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func Up(network *net.Network) (string, error) {
	url := "/v1/server/up"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Version(network *net.Network) (string, error) {
	url := "/v1/server/version"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetConfig(network *net.Network, configKey string) (string, error) {
	url := fmt.Sprintf("/v1/server/config/%s", configKey)
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Reload(network *net.Network) (string, error) {
	url := "/v1/server/properties/reload"
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Status(network *net.Network) (string, error) {
	url := "/v1/server/status"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func User(network *net.Network) (string, error) {
	url := "/v1/server/user"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func ShuttingDown(network *net.Network) (string, error) {
	url := "/v1/server/shuttingDown"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Healthy(network *net.Network) (string, error) {
	url := "/v1/server/healthy"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func UpExtended(network *net.Network) (string, error) {
	url := "/v1/server/up/extended"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func State(network *net.Network) (string, error) {
	url := "/v1/server/ha/state"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// WIP
func SetState(network *net.Network) (string, error) {
	url := "/v1/server/ha/state"
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Metrics(network *net.Network) (string, error) {
	url := "/v1/server/ha/metrics"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func Priority(network *net.Network) (string, error) {
	url := "/v1/server/ha/priority"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// WIP
func SetPriority(network *net.Network) (string, error) {
	url := "/v1/server/ha/priority"
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func States(network *net.Network) (string, error) {
	url := "/v1/server/ha/states"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", nil
	}
	return string(body)
}

// WIP
func ClearStates(network *net.Network) (string, error) {
	url := "/v1/server/ha/states/clear"
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		return "", nil
	}
	return string(body)
}

func Export(network *net.Network) (string, error) {
	url := "/v1/server/ha/persist/export"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", nil
	}
	return string(body)
}

// WIP
func Shutdown(network *net.Network) (string, error) {
	url := "/v1/server/shutdown"
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		return "", err
	}
	return string(body)
}
