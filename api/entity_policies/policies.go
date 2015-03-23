package entity_policies

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func PolicyList(application, entity string) []models.PolicySummary {
	url := "http://192.168.50.101:8081/v1/applications/" + application + "/entities/"+ entity + "/policies"
	req := net.NewGetRequest(url)
	body, err := net.SendRequest(req)
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

func PolicyStatus(application, entity, policy string) string {
	url := "http://192.168.50.101:8081/v1/applications/" + application + "/entities/"+ entity + "/policies/" + policy
	req := net.NewGetRequest(url)
	body, err := net.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}