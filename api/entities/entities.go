package entities

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"net/url"
)

//WIP
func GetTask(network *net.Network, application, entity, task string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/activities/%s", application, entity, task)
	body, err := network.SendGetRequest(url)
    // TODO return model
	return string(body), err
}

//WIP
func GetIcon(network *net.Network, application, entity string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/icon", application, entity)
	body, err := network.SendGetRequest(url)
    // TODO return model
	return string(body), err
}

func Children(network *net.Network, application, entity string) ([]models.EntitySummary, error) {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities/%s/children", application, entity)
    var entityList []models.EntitySummary
    body, err := network.SendGetRequest(urlStr)
    if err != nil {
        return entityList, err
    }

	err = json.Unmarshal(body, &entityList)
	return entityList, err
}

func AddChildren(network *net.Network, application, entity, filePath string) (models.TaskSummary, error) {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities/%s/children", application, entity)
    var response models.TaskSummary
    body, err := network.SendPostFileRequest(urlStr, filePath, "application/json")
    if err != nil {
        return response, err
    }

	err = json.Unmarshal(body, &response)
	return response, err
}

//WIP
func GetLocations(network *net.Network, application, entity string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/locations", application, entity)
	body, err := network.SendGetRequest(url)
    // TODO return model
	return string(body), err
}

func Spec(network *net.Network, application, entity string) (string, error) {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities/%s/spec", application, entity)
	body, err := network.SendGetRequest(urlStr)
	return string(body), err
}

//WIP
func GetDescendants(network *net.Network, application, entity string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/descendants", application, entity)
	body, err := network.SendGetRequest(url)
    // TODO return model
	return string(body), err
}

//WIP
func GetDescendantsSensor(network *net.Network, application, entity, sensor string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/descendants/sensor/%s", application, entity, sensor)
	body, err := network.SendGetRequest(url)
    // TODO return model
	return string(body), err
}

func GetActivities(network *net.Network, application, entity string) ([]models.TaskSummary, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/activities", application, entity)
    var activityList []models.TaskSummary
    body, err := network.SendGetRequest(url)
    if err != nil {
        return activityList, err
    }

	err = json.Unmarshal(body, &activityList)
	return activityList, err
}

//WIP
func GetTags(network *net.Network, application, entity string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/tags", application, entity)
	body, err := network.SendGetRequest(url)
    // TODO return model
	return string(body), err
}

//WIP
func Expunge(network *net.Network, application, entity string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/expunge", application, entity)
	body, err := network.SendEmptyPostRequest(url)
    // TODO return model
	return string(body), err
}

//WIP
func GetEntity(network *net.Network, application, entity string) (models.EntitySummary, error) {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s", application, entity)
    summary := models.EntitySummary{}
    body, err := network.SendGetRequest(url)
	if err != nil {
		return summary, err
	}

    err = json.Unmarshal(body, &summary)
    return summary, err
}

func EntityList(network *net.Network, application string) ([]models.EntitySummary, error) {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities", application)
    var entityList []models.EntitySummary
    body, err := network.SendGetRequest(urlStr)
    if err != nil {
        return entityList, err
    }

	err = json.Unmarshal(body, &entityList)
	return entityList, err
}

func Rename(network *net.Network, application, entity, newName string) (string, error) {
	urlStr := fmt.Sprintf("/v1/applications/%s/entities/%s/name?name=%s", application, entity, url.QueryEscape(newName))
	body, err := network.SendEmptyPostRequest(urlStr)
	return string(body), err
}
