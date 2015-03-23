package entity_policies

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func PolicyList(network *net.Network, application, entity string) []models.PolicySummary {
	url := "/v1/applications/" + application + "/entities/"+ entity + "/policies"
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
	url := "/v1/applications/" + application + "/entities/"+ entity + "/policies/" + policy
	req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}