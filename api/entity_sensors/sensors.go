package entity_sensors

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func SensorList(application, entity string) []models.SensorSummary {
	url := "http://192.168.50.101:8081/v1/applications/" + application + "/entities/"+ entity + "/sensors"
	req := net.NewGetRequest(url)
	body, err := net.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	
	var sensorList []models.SensorSummary
	err = json.Unmarshal(body, &sensorList)
	if err != nil{
		fmt.Println(err)
	}
	return sensorList
}

func SensorValue(application, entity, sensor string) string {
	url := "http://192.168.50.101:8081/v1/applications/" + application + "/entities/"+ entity + "/sensors/" + sensor
	req := net.NewGetRequest(url)
	body, err := net.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}

	return string(body)
}