package entity_sensors

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)



func SensorValue(network *net.Network, application, entity, sensor string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors/%s", application, entity, sensor)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}

// WIP
func DeleteSensor(network *net.Network, application, entity, sensor string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors/%s", application, entity, sensor)
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}

// WIP
//func SetSensor(network *net.Network, application, entity, sensor string) string {
//	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors/%s", application, entity, sensor)
//	body, err := network.SendPostRequest(url)
//	if err != nil {
//		fmt.Println(err)
//	}

//	return string(body)
//}

// WIP
//func SetSensors(network *net.Network, application, entity, sensor string) string {
//	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors", application, entity, sensor)
//	body, err := network.SendPostRequest(url)
//	if err != nil {
//		fmt.Println(err)
//	}

//	return string(body)
//}

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


func CurrentState(network *net.Network, application, entity, sensor string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors/current-state", application, entity, sensor)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}
