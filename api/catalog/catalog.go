package catalog

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func Icon(network *net.Network, itemId string) []byte {
	url := fmt.Sprintf("/v1/catalog/icon/%s", itemId)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

func IconWithVersion(network *net.Network, itemId, version string) []byte {
	url := fmt.Sprintf("/v1/catalog/icon/%s/%s", itemId, version)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

func GetEntityWithVersion(network *net.Network, entityId, version string) models.CatalogEntitySummary {
	url := fmt.Sprintf("/v1/catalog/entities/%s/%s", entityId, version)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var catalogEntity models.CatalogEntitySummary
	err = json.Unmarshal(body, &catalogEntity)
	return catalogEntity
}

func DeleteEntityWithVersion(network *net.Network, entityId, version string) string {
	url := fmt.Sprintf("/v1/catalog/entities/%s/%s", entityId, version)
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func GetEntity(network *net.Network, entityId string) models.CatalogEntitySummary {
	url := fmt.Sprintf("/v1/catalog/entities/%s", entityId)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var catalogEntity models.CatalogEntitySummary
	err = json.Unmarshal(body, &catalogEntity)
	return catalogEntity
}

func DeleteEntity(network *net.Network, entityId string) string {
	url := fmt.Sprintf("/v1/catalog/entities/%s", entityId)
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func GetPolicy(network *net.Network, policyId string) models.CatalogItemSummary {
	url := fmt.Sprintf("/v1/catalog/policies/%s", policyId)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var catalogItem models.CatalogItemSummary
	err = json.Unmarshal(body, &catalogItem)
	return catalogItem
}

func GetPolicyWithVersion(network *net.Network, policyId, version string) models.CatalogItemSummary {
	url := fmt.Sprintf("/v1/catalog/policies/%s/%s", policyId)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var catalogItem models.CatalogItemSummary
	err = json.Unmarshal(body, &catalogItem)
	return catalogItem
}

func DeletePolicyWithVersion(network *net.Network, policyId, version string) string {
	url := fmt.Sprintf("/v1/catalog/policies/%s/%s", policyId)
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func GetApplication(network *net.Network, applicationId string) models.CatalogEntitySummary {
	url := fmt.Sprintf("/v1/catalog/applications/%s", applicationId)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var catalogEntity models.CatalogEntitySummary
	err = json.Unmarshal(body, &catalogEntity)
	return catalogEntity
}

func GetApplicationWithVersion(network *net.Network, applicationId, version string) models.CatalogEntitySummary {
	url := fmt.Sprintf("/v1/catalog/applications/%s/%s", applicationId, version)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var catalogEntity models.CatalogEntitySummary
	err = json.Unmarshal(body, &catalogEntity)
	return catalogEntity
}

func DeleteApplicationWithVersion(network *net.Network, applicationId, version string) string {
	url := fmt.Sprintf("/v1/catalog/applications/%s/%s", applicationId, version)
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
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

func Locations(network *net.Network) models.CatalogLocationSummary {
	url := "/v1/catalog/locations"
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var catalogLocation models.CatalogLocationSummary
	err = json.Unmarshal(body, &catalogLocation)
	return catalogLocation
}

func AddCatalog(network *net.Network, filePath string) string {
	url := "/v1/catalog"
	body, err := network.SendPostFileRequest(url, filePath)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func Reset(network *net.Network) string {
	url := "/v1/catalog/reset"
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func GetLocationWithVersion(network *net.Network, locationId, version string) models.CatalogLocationSummary {
	url := fmt.Sprintf("/v1/catalog/locations/%s/%s", locationId, version)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var catalogLocation models.CatalogLocationSummary
	err = json.Unmarshal(body, &catalogLocation)
	return catalogLocation
}

func PostLocationWithVersion(network *net.Network, locationId, version string) string  {
	url := fmt.Sprintf("/v1/catalog/locations/%s/%s", locationId, version)
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func Entities(network *net.Network) []models.CatalogItemSummary {
	url := "/v1/catalog/entities"
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var applications []models.CatalogItemSummary
	err = json.Unmarshal(body, &applications)
	return applications
}

func Catalog(network *net.Network) []models.CatalogItemSummary {
	url := "/v1/catalog/applications"
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var applications []models.CatalogItemSummary
	err = json.Unmarshal(body, &applications)
	return applications
}

func GetLocation(network *net.Network, locationId string) models.CatalogLocationSummary {
	url := fmt.Sprintf("/v1/catalog/locations/%s", locationId)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var catalogLocation models.CatalogLocationSummary
	err = json.Unmarshal(body, &catalogLocation)
	return catalogLocation
}
