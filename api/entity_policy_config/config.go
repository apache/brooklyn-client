package entity_policies_config

import (
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"net/url"
)

func CurrentState(network *net.Network, application, entity, policy string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/config/current-state", application, entity, policy)
	body, err := network.SendGetRequest(url)
    if nil != err {
        return "", err
    }
	return string(body), nil
}

func GetConfigValue(network *net.Network, application, entity, policy, config string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/config/%s", application, entity, policy, config)
	body, err := network.SendGetRequest(url)
    if nil != err {
        return "", err
    }
	return string(body), nil
}

// WIP
func SetConfigValue(network *net.Network, application, entity, policy, config string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/config/%s", application, entity, policy, config)
	body, err := network.SendEmptyPostRequest(url)
    if nil != err {
        return "", err
    }
	return string(body), nil
}

func GetAllConfigValues(network *net.Network, application, entity, policy, config string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/config", application, entity, policy, config)
	body, err := network.SendGetRequest(url)
    if nil != err {
        return "", err
    }
	return string(body), nil
}