package entities

import (
	"encoding/json"
	"fmt"
	"github.com/robertgmoss/brooklyn-cli/models"
	"github.com/robertgmoss/brooklyn-cli/net"
	"os"
	"path/filepath"
)

func Spec(network *net.Network, application, entity string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/spec", application, entity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func EntityList(network *net.Network, application string) []models.EntitySummary {
	url := fmt.Sprintf("/v1/applications/%s/entities", application)
	body, err := network.SendGetRequest(url)
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

func Children(network *net.Network, application, entity string) []models.EntitySummary {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/children", application, entity)
	body, err := network.SendGetRequest(url)
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
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/children", application, entity)
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	req := network.NewPostRequest(url, file)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	body, err := network.SendRequest(req)
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
