package catalog

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func Icon(network *net.Network, itemId string) ([]byte, error) {
	url := fmt.Sprintf("/v1/catalog/icon/%s", itemId)
	body, err := network.SendGetRequest(url)
	return body, err
}

func IconWithVersion(network *net.Network, itemId, version string) ([]byte, error) {
	url := fmt.Sprintf("/v1/catalog/icon/%s/%s", itemId, version)
	body, err := network.SendGetRequest(url)
	return body, err
}

func GetEntityWithVersion(network *net.Network, entityId, version string) (models.CatalogEntitySummary, error) {
	url := fmt.Sprintf("/v1/catalog/entities/%s/%s", entityId, version)
    var catalogEntity models.CatalogEntitySummary
    body, err := network.SendGetRequest(url)
    if err != nil {
        return catalogEntity, err
    }
	err = json.Unmarshal(body, &catalogEntity)
	return catalogEntity, err
}

func DeleteEntityWithVersion(network *net.Network, entityId, version string) (string, error) {
	url := fmt.Sprintf("/v1/catalog/entities/%s/%s", entityId, version)
	body, err := network.SendDeleteRequest(url)
	return string(body), err
}

func GetEntity(network *net.Network, entityId string) (models.CatalogEntitySummary, error) {
	url := fmt.Sprintf("/v1/catalog/entities/%s", entityId)
    var catalogEntity models.CatalogEntitySummary
    body, err := network.SendGetRequest(url)
    if err != nil {
        return catalogEntity, err
    }
	err = json.Unmarshal(body, &catalogEntity)
	return catalogEntity, err
}

func DeleteEntity(network *net.Network, entityId string) (string, error) {
	url := fmt.Sprintf("/v1/catalog/entities/%s", entityId)
	body, err := network.SendDeleteRequest(url)
	return string(body), err
}

func GetPolicy(network *net.Network, policyId string) (models.CatalogItemSummary, error) {
	url := fmt.Sprintf("/v1/catalog/policies/%s", policyId)
    var catalogItem models.CatalogItemSummary
    body, err := network.SendGetRequest(url)
    if err != nil {
        return catalogItem, err
    }
	err = json.Unmarshal(body, &catalogItem)
	return catalogItem, err
}

func GetPolicyWithVersion(network *net.Network, policyId, version string) (models.CatalogItemSummary, error) {
	url := fmt.Sprintf("/v1/catalog/policies/%s/%s", policyId)
    var catalogItem models.CatalogItemSummary
    body, err := network.SendGetRequest(url)
    if err != nil {
        return catalogItem, err
    }
	err = json.Unmarshal(body, &catalogItem)
	return catalogItem, err
}

func DeletePolicyWithVersion(network *net.Network, policyId, version string) (string, error) {
	url := fmt.Sprintf("/v1/catalog/policies/%s/%s", policyId)
	body, err := network.SendDeleteRequest(url)
	return string(body), err
}

func GetApplication(network *net.Network, applicationId string) (models.CatalogEntitySummary, error) {
	url := fmt.Sprintf("/v1/catalog/applications/%s", applicationId)
    var catalogEntity models.CatalogEntitySummary
    body, err := network.SendGetRequest(url)
    if err != nil {
        return catalogEntity, err
    }
	err = json.Unmarshal(body, &catalogEntity)
	return catalogEntity, err
}

func GetApplicationWithVersion(network *net.Network, applicationId, version string) (models.CatalogEntitySummary, error) {
	url := fmt.Sprintf("/v1/catalog/applications/%s/%s", applicationId, version)
    var catalogEntity models.CatalogEntitySummary
    body, err := network.SendGetRequest(url)
    if err != nil {
        return catalogEntity, err
    }
	err = json.Unmarshal(body, &catalogEntity)
	return catalogEntity, err
}

func DeleteApplicationWithVersion(network *net.Network, applicationId, version string) (string, error) {
	url := fmt.Sprintf("/v1/catalog/applications/%s/%s", applicationId, version)
	body, err := network.SendDeleteRequest(url)
	return string(body), err
}

func Policies(network *net.Network) ([]models.CatalogPolicySummary, error) {
	url := "/v1/catalog/policies"
    var policies []models.CatalogPolicySummary
    body, err := network.SendGetRequest(url)
    if err != nil {
        return policies, err
    }
	err = json.Unmarshal(body, &policies)
	return policies, err
}

func Locations(network *net.Network) (models.CatalogLocationSummary, error) {
	url := "/v1/catalog/locations"
    var catalogLocation models.CatalogLocationSummary
    body, err := network.SendGetRequest(url)
    if err != nil {
        return catalogLocation, err
    }
	err = json.Unmarshal(body, &catalogLocation)
	return catalogLocation, err
}

func AddCatalog(network *net.Network, filePath string) (string, error) {
	url := "/v1/catalog"
	body, err := network.SendPostFileRequest(url, filePath)
	return string(body), err
}

func Reset(network *net.Network) (string, error) {
	url := "/v1/catalog/reset"
	body, err := network.SendEmptyPostRequest(url)
	return string(body), err
}

func GetLocationWithVersion(network *net.Network, locationId, version string) (models.CatalogLocationSummary, error) {
	url := fmt.Sprintf("/v1/catalog/locations/%s/%s", locationId, version)
    var catalogLocation models.CatalogLocationSummary
    body, err := network.SendGetRequest(url)
    if err != nil {
        return catalogLocation, err
    }
	err = json.Unmarshal(body, &catalogLocation)
	return catalogLocation, err
}

func PostLocationWithVersion(network *net.Network, locationId, version string) (string, error)  {
	url := fmt.Sprintf("/v1/catalog/locations/%s/%s", locationId, version)
	body, err := network.SendEmptyPostRequest(url)
	return string(body), err
}

func Entities(network *net.Network) ([]models.CatalogItemSummary, error) {
	url := "/v1/catalog/entities"
    var entities []models.CatalogItemSummary
    body, err := network.SendGetRequest(url)
    if err != nil {
        return entities, err
    }
	err = json.Unmarshal(body, &entities)
	return entities, err
}

func Catalog(network *net.Network) ([]models.CatalogItemSummary, error) {
	url := "/v1/catalog/applications"
    var applications []models.CatalogItemSummary
    body, err := network.SendGetRequest(url)
    if err != nil {
        return applications, err
    }
	err = json.Unmarshal(body, &applications)
	return applications, err
}

func GetLocation(network *net.Network, locationId string) (models.CatalogLocationSummary, error) {
	url := fmt.Sprintf("/v1/catalog/locations/%s", locationId)
    var catalogLocation models.CatalogLocationSummary
    body, err := network.SendGetRequest(url)
    if err != nil {
        return catalogLocation, err
    }
	err = json.Unmarshal(body, &catalogLocation)
	return catalogLocation, err
}
