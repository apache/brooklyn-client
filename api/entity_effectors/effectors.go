package entity_effectors

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func EffectorList(network *net.Network, application, entity string) []models.EffectorSummary {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/effectors", application, entity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var effectorList []models.EffectorSummary
	err = json.Unmarshal(body, &effectorList)
	if err != nil {
		fmt.Println(err)
	}
	return effectorList
}

func TriggerEffector(network *net.Network, application, entity, effector string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/effectors/%s", application, entity, effector)
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}
