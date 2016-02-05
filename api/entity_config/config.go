package entity_config

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func ConfigValue(network *net.Network, application, entity, config string) (string, error) {
	bytes, err := ConfigValueAsBytes(network, application, entity, config)
	if nil != err {
		return "", err
	}

	var value string;
	err = json.Unmarshal(bytes, &value)
	if nil == err {
		return value, nil
	}
	return string(bytes), nil
}

func ConfigValueAsBytes(network *net.Network, application, entity, config string) ([]byte, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config/%s", application, entity, config)
	body, err := network.SendGetRequest(url)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func SetConfig(network *net.Network, application, entity, config, value string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config/%s", application, entity, config)
	val := []byte(value)
	body, err := network.SendPostRequest(url, val)
    if nil != err {
        return "", err
    }
    return string(body), nil
}

func ConfigList(network *net.Network, application, entity string) ([]models.ConfigSummary, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config", application, entity)
    var configList []models.ConfigSummary
    body, err := network.SendGetRequest(url)
    if err != nil {
        return configList, err
    }

	err = json.Unmarshal(body, &configList)
	return configList, err
}

func PostConfig(network *net.Network, application, entity, config, value string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config", application, entity)
	val := []byte(value)
	body, err := network.SendPostRequest(url, val)
    if nil != err {
        return "", err
    }
    return string(body), nil
}



func ConfigCurrentState(network *net.Network, application, entity string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config/current-state", application, entity)
    var currentState map[string]interface{}
    body, err := network.SendGetRequest(url)
    if err != nil {
        return currentState, err
    }
	err = json.Unmarshal(body, &currentState)
	return currentState, err
}
