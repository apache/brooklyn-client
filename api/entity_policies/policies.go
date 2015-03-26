package entity_policies

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/robertgmoss/brooklyn-cli/models"
	"github.com/robertgmoss/brooklyn-cli/net"
	"net/url"
	"strconv"
)

// WIP
func AddPolicy(network *net.Network, application, entity, policy string, config map[string]string) models.PolicySummary {
	path := fmt.Sprintf("/v1/applications/%s/entities/%s/policies", application, entity)
	data := url.Values{}
	data.Set("policyType", policy)
	//data.Add("config", config)
	req := network.NewPostRequest(path, bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	var policySummary models.PolicySummary
	err = json.Unmarshal(body, &policySummary)
	if err != nil {
		fmt.Println(err)
	}
	return policySummary
}

func StartPolicy(network *net.Network, application, entity, policy string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/start", application, entity, policy)
	req := network.NewPostRequest(url, nil)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func StopPolicy(network *net.Network, application, entity, policy string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/stop", application, entity, policy)
	req := network.NewPostRequest(url, nil)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func DestroyPolicy(network *net.Network, application, entity, policy string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s/destroy", application, entity, policy)
	req := network.NewPostRequest(url, nil)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func PolicyList(network *net.Network, application, entity string) []models.PolicySummary {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies", application, entity)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var policyList []models.PolicySummary
	err = json.Unmarshal(body, &policyList)
	if err != nil {
		fmt.Println(err)
	}
	return policyList
}

func PolicyStatus(network *net.Network, application, entity, policy string) string {
	url := fmt.Sprintf("/v1/applications/%s/entities/%s/policies/%s", application, entity, policy)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}
