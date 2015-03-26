package entity_sensors

import (
	"encoding/json"
	"fmt"
	"github.com/robertgmoss/brooklyn-cli/models"
	"github.com/robertgmoss/brooklyn-cli/net"
)

func SensorList(network *net.Network, application, entity string) []models.SensorSummary {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors", application, entity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var sensorList []models.SensorSummary
	err = json.Unmarshal(body, &sensorList)
	if err != nil {
		fmt.Println(err)
	}
	return sensorList
}

func SensorValue(network *net.Network, application, entity, sensor string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors/%s", application, entity, sensor)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}
