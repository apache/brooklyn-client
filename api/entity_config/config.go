package entity_config

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func ConfigList(network *net.Network, application, entity string) []models.ConfigSummary {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config", application, entity)
    req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	
	var configList []models.ConfigSummary
	err = json.Unmarshal(body, &configList)
	if err != nil{
		fmt.Println(err)
	}
	return configList
}

func ConfigValue(network *net.Network, application, entity, config string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config/%s", application, entity, config)
	req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}

func ConfigCurrentState(network *net.Network, application, entity string) map[string]interface{} {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/config/current-state", application, entity)
    req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	
	var currentState map[string]interface{}
	err = json.Unmarshal(body, &currentState)
	if err != nil{
		fmt.Println(err)
	}
	return currentState
}