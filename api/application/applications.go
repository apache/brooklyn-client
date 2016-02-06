package application

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

//WIP
func Fetch(network *net.Network) (string, error) {
	url := "/v1/applications/fetch"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	// TODO return model
	return string(body), nil
}

func Applications(network *net.Network) ([]models.ApplicationSummary, error) {
	url := fmt.Sprintf("/v1/applications")
	var appSummary []models.ApplicationSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return appSummary, err
	}

	err = json.Unmarshal(body, &appSummary)
	return appSummary, err
}

func Create(network *net.Network, filePath string) (models.TaskSummary, error) {
	url := "/v1/applications"
	var response models.TaskSummary
	body, err := network.SendPostFileRequest(url, filePath, "application/json")
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}

func CreateFromBytes(network *net.Network, blueprint []byte) (models.TaskSummary, error) {
	url := "/v1/applications"
	var response models.TaskSummary
	body, err := network.SendPostRequest(url, blueprint)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}

// WIP
func Descendants(network *net.Network, app string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/descendants", app)

	body, err := network.SendGetRequest(url)
	// TODO return model
	if nil != err {
		return "", err
	}
	return string(body), nil
}

// WIP
func DescendantsSensor(network *net.Network, app, sensor string) (string, error) {
	url := fmt.Sprintf("/v1/applications/%s/descendants/sensor/%s", app, sensor)

	body, err := network.SendGetRequest(url)
	// TODO return model
	if nil != err {
		return "", err
	}
	return string(body), nil
}

func Tree(network *net.Network) ([]models.Tree, error) {
	url := "/v1/applications/tree"
	var tree []models.Tree
	body, err := network.SendGetRequest(url)
	if err != nil {
		return tree, err
	}

	err = json.Unmarshal(body, &tree)
	return tree, err
}

func Application(network *net.Network, app string) (models.ApplicationSummary, error) {
	url := fmt.Sprintf("/v1/applications/%s", app)
	var appSummary models.ApplicationSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return appSummary, err
	}

	err = json.Unmarshal(body, &appSummary)
	return appSummary, err
}

func Delete(network *net.Network, application string) (models.TaskSummary, error) {
	url := fmt.Sprintf("/v1/applications/%s", application)
	var response models.TaskSummary
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response)
	return response, err
}

// WIP
func CreateLegacy(network *net.Network) (string, error) {
	url := fmt.Sprintf("/v1/applications/createLegacy")
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		return "", err
	}
	// TODO return model
	return string(body), nil
}
