package entity_config

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func ConfigValue(network *net.Network, application, entity, config string) string {
	return string(ConfigValueAsBytes(network, application, entity, config))
}

func ConfigValueAsBytes(network *net.Network, application, entity, config string) []byte {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config/%s", application, entity, config)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	return body
}

func SetConfig(network *net.Network, application, entity, config, value string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config/%s", application, entity, config)
	val := []byte(value)
	body, err := network.SendPostRequest(url, val)
	if err != nil {
		fmt.Println(err)
	}
	
	return string(body)
}

func ConfigList(network *net.Network, application, entity string) []models.ConfigSummary {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config", application, entity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var configList []models.ConfigSummary
	err = json.Unmarshal(body, &configList)
	if err != nil {
		fmt.Println(err)
	}
	return configList
}

func PostConfig(network *net.Network, application, entity, config, value string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config", application, entity)
	val := []byte(value)
	body, err := network.SendPostRequest(url, val)
	if err != nil {
		fmt.Println(err)
	}
	
	return string(body)
}



func ConfigCurrentState(network *net.Network, application, entity string) map[string]interface{} {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config/current-state", application, entity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var currentState map[string]interface{}
	err = json.Unmarshal(body, &currentState)
	if err != nil {
		fmt.Println(err)
	}
	return currentState
}
