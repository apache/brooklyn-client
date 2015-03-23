package entity_policies

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func PolicyList(network *net.Network, application, entity string) []models.PolicySummary {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies", application, entity)
    req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	
	var policyList []models.PolicySummary
	err = json.Unmarshal(body, &policyList)
	if err != nil{
		fmt.Println(err)
	}
	return policyList
}

func PolicyStatus(network *net.Network, application, entity, policy string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s", application, entity, policy)
    req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}