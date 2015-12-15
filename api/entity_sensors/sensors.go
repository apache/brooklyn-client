package entity_sensors

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)



func SensorValue(network *net.Network, application, entity, sensor string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors/%s", application, entity, sensor)
	body, err := network.SendGetRequest(url)
    if nil != err {
        return "", err
    }
	return string(body), nil
}

// WIP
func DeleteSensor(network *net.Network, application, entity, sensor string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors/%s", application, entity, sensor)
	body, err := network.SendDeleteRequest(url)
    if nil != err {
        return "", err
    }
	return string(body), nil
}

// WIP
//func SetSensor(network *net.Network, application, entity, sensor string) string {
//	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors/%s", application, entity, sensor)
//	body, err := network.SendPostRequest(url)
//	if err != nil {
//		error_handler.ErrorExit(err)
//	}

//	return string(body)
//}

// WIP
//func SetSensors(network *net.Network, application, entity, sensor string) string {
//	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors", application, entity, sensor)
//	body, err := network.SendPostRequest(url)
//	if err != nil {
//		error_handler.ErrorExit(err)
//	}

//	return string(body)
//}

func SensorList(network *net.Network, application, entity string) ([]models.SensorSummary, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors", application, entity)
	body, err := network.SendGetRequest(url)
    var sensorList []models.SensorSummary
    if err != nil {
        return sensorList, err
    }

	err = json.Unmarshal(body, &sensorList)
	return sensorList, err
}


func CurrentState(network *net.Network, application, entity string) (map[string]interface{}, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/sensors/current-state", application, entity)
    var currentState map[string]interface{}
    body, err := network.SendGetRequest(url)
    if err != nil {
        return currentState, err
    }

    err = json.Unmarshal(body, &currentState)
	return currentState, err
}
