package activities

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func Activity(network *net.Network, activity string) (models.TaskSummary, error) {
	url := fmt.Sprintf("/v1/activities/%s", activity)
	var task models.TaskSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return task, err
	}

	err = json.Unmarshal(body, &task)
	return task, err
}

func ActivityChildren(network *net.Network, activity string) ([]models.TaskSummary, error) {
	url := fmt.Sprintf("/v1/activities/%s/children", activity)
	var tasks []models.TaskSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return tasks, err
	}

	err = json.Unmarshal(body, &tasks)
	return tasks, err
}

func ActivityStream(network *net.Network, activity, streamId string) (string, error) {
	url := fmt.Sprintf("/v1/activities/%s/stream/%s", activity, streamId)
	body, err := network.SendGetRequest(url)
	if nil != err {
		return "", err
	}
	return string(body), nil
}
