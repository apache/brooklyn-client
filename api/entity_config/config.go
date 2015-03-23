package entity_config

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func ConfigList(application, entity string) []models.ConfigSummary {
	url := "http://192.168.50.101:8081/v1/applications/" + application + "/entities/"+ entity + "/config"
	req := net.NewGetRequest(url)
	body, err := net.SendRequest(req)
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

func ConfigValue(application, entity, config string) string {
	url := "http://192.168.50.101:8081/v1/applications/" + application + "/entities/"+ entity + "/config/" + config
	req := net.NewGetRequest(url)
	body, err := net.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}