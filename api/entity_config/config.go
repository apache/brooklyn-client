package entity_config

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func ConfigList(network *net.Network, application, entity string) []models.ConfigSummary {
	url := "/v1/applications/" + application + "/entities/"+ entity + "/config"
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
	url := "/v1/applications/" + application + "/entities/"+ entity + "/config/" + config
	req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}