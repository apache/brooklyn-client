package activities

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func Activity(network *net.Network, activity string) models.TaskSummary {
	url := fmt.Sprintf("/v1/activities/%s", activity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var task models.TaskSummary
	err = json.Unmarshal(body, &task)
	if err != nil {
		fmt.Println(err)
	}
	return task
}

func ActivityChildren(network *net.Network, activity string) []models.TaskSummary {
	url := fmt.Sprintf("/v1/activities/%s/children", activity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var tasks []models.TaskSummary
	err = json.Unmarshal(body, &tasks)
	if err != nil {
		fmt.Println(err)
	}
	return tasks
}

func ActivityStream(network *net.Network, activity, streamId string) string {
	url := fmt.Sprintf("/v1/activities/%s/stream/%s", activity, streamId)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}
