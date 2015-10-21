package entities

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"net/url"
)

//WIP
func GetTask(network *net.Network, application, entity, task string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/activities/%s", application, entity, task)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
    // TODO return model
	return string(body)
}

//WIP
func GetIcon(network *net.Network, application, entity string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/icon", application, entity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
    // TODO return model
	return string(body)
}

func Children(network *net.Network, application, entity string) []models.EntitySummary {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities/%s/children", application, entity)
	body, err := network.SendGetRequest(urlStr)
	if err != nil {
		fmt.Println(err)
	}

	var entityList []models.EntitySummary
	err = json.Unmarshal(body, &entityList)
	if err != nil {
		fmt.Println(err)
	}
	return entityList
}

func AddChildren(network *net.Network, application, entity, filePath string) models.TaskSummary {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities/%s/children", application, entity)
	body, err := network.SendPostFileRequest(urlStr, filePath)
	if err != nil {
		fmt.Println(err)
	}

	var response models.TaskSummary
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err)
	}
	return response
}

//WIP
func GetLocations(network *net.Network, application, entity string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/locations", application, entity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
    // TODO return model
	return string(body)
}

func Spec(network *net.Network, application, entity string) string {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities/%s/spec", application, entity)
	body, err := network.SendGetRequest(urlStr)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

//WIP
func GetDescendants(network *net.Network, application, entity string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/descendants", application, entity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
    // TODO return model
	return string(body)
}

//WIP
func GetDescendantsSensor(network *net.Network, application, entity, sensor string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/descendants/sensor/%s", application, entity, sensor)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
    // TODO return model
	return string(body)
}

func GetActivities(network *net.Network, application, entity string) []models.TaskSummary {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/activities", application, entity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var activityList []models.TaskSummary
	err = json.Unmarshal(body, &activityList)
	if err != nil {
		fmt.Println(err)
	}
	return activityList
}

//WIP
func GetTags(network *net.Network, application, entity string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/tags", application, entity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
    // TODO return model
	return string(body)
}

//WIP
func Expunge(network *net.Network, application, entity string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/expunge", application, entity)
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		fmt.Println(err)
	}
    // TODO return model
	return string(body)
}

//WIP
func GetEntity(network *net.Network, application, entity string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s", application, entity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
    // TODO return model
	return string(body)
}

func EntityList(network *net.Network, application string) []models.EntitySummary {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities", application)
	body, err := network.SendGetRequest(urlStr)
	if err != nil {
		fmt.Println(err)
	}

	var entityList []models.EntitySummary
	err = json.Unmarshal(body, &entityList)
	if err != nil {
		fmt.Println(err)
	}
	return entityList
}

func Rename(network *net.Network, application, entity, newName string) string {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities/%s/name?name=%s", application, entity, url.QueryEscape(newName))
	body, err := network.SendEmptyPostRequest(urlStr)
	if err != nil {
		fmt.Println(err)
	}
	
	return string(body)
}
