package application

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

//WIP
func Fetch(network *net.Network) string {
	url := "/v1/applications/fetch"
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	// TODO return model
	return string(body)
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
	body, err := network.SendPostFileRequest(url, filePath, "application/json")
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

func CreateFromBytes(network *net.Network, blueprint []byte)  models.TaskSummary {
	url := "/v1/applications"
	
	body, err := network.SendPostRequest(url, blueprint)
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

// WIP
func Descendants(network *net.Network, app string) string {
	url := fmt.Sprintf("/v1/applications/%s/descendants", app) 
	
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	// TODO return model
	return string(body)
}

// WIP
func DescendantsSensor(network *net.Network, app, sensor string) string {
	url := fmt.Sprintf("/v1/applications/%s/descendants/sensor/%s", app, sensor) 
	
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	// TODO return model
	return string(body)
}

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

// WIP
func CreateLegacy(network *net.Network) string {
	url := fmt.Sprintf("/v1/applications/createLegacy")
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	// TODO return model
	return string(body)
}


