package catalog

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func Catalog(network *net.Network) []models.Application {
	url := "/v1/catalog/applications"
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var applications []models.Application
	err = json.Unmarshal(body, &applications)
	return applications
}

func Entities(network *net.Network) []models.Application {
	url := "/v1/catalog/entities"
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var applications []models.Application
	err = json.Unmarshal(body, &applications)
	return applications
}

func Policies(network *net.Network) []models.CatalogPolicySummary {
	url := "/v1/catalog/policies"
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var policies []models.CatalogPolicySummary
	err = json.Unmarshal(body, &policies)
	return policies
}

func AddCatalog(network *net.Network, filePath string) string {
	url := "/v1/catalog"
	body, err := network.SendPostFileRequest(url, filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}
