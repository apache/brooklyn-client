package application

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func Tree(network *net.Network) []models.Tree {
	url := "/v1/applications/tree"
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var tree []models.Tree
	err = json.Unmarshal(body, &tree)
	if err != nil {
		fmt.Println(err)
	}
	return tree
}

func Application(network *net.Network, app string) models.ApplicationSummary {
	url := fmt.Sprintf("/v1/applications/%s", app)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var appSummary models.ApplicationSummary
	err = json.Unmarshal(body, &appSummary)
	if err != nil {
		fmt.Println(err)
	}
	return appSummary
}

func Applications(network *net.Network) []models.ApplicationSummary {
	url := fmt.Sprintf("/v1/applications")
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var appSummary []models.ApplicationSummary
	err = json.Unmarshal(body, &appSummary)
	if err != nil {
		fmt.Println(err)
	}
	return appSummary
}

func Create(network *net.Network, filePath string) models.TaskSummary {
	url := "/v1/applications"
	body, err := network.SendPostFileRequest(url, filePath)
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

func Delete(network *net.Network, application string) models.TaskSummary {
	url := fmt.Sprintf("/v1/applications/%s", application)
	body, err := network.SendDeleteRequest(url)
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
