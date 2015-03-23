package entities

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func EntityList(network *net.Network, application string) []models.EntitySummary {
	url := fmt.Sprintf("/v1/applications/%s/entities", application)
    req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	
	var entityList []models.EntitySummary
	err = json.Unmarshal(body, &entityList)
	if err != nil{
		fmt.Println(err)
	}
	return entityList
}

func ActivityList(network *net.Network, application, entity string) []models.TaskSummary {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/activities", application, entity)
    req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	
	var activityList []models.TaskSummary
	err = json.Unmarshal(body, &activityList)
	if err != nil{
		fmt.Println(err)
	}
	return activityList
}