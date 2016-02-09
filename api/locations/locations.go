package locations

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func LocatedLocations(network *net.Network) (string, error) {
	url := "/v1/locations/usage/LocatedLocations"
	body, err := network.SendGetRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetLocation(network *net.Network, locationId string) (models.LocationSummary, error) {
	url := fmt.Sprintf("/v1/locations/%s", locationId)
	var locationDetail models.LocationSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return locationDetail, err
	}
	err = json.Unmarshal(body, &locationDetail)
	return locationDetail, err
}

func DeleteLocation(network *net.Network, locationId string) (string, error) {
	url := fmt.Sprintf("/v1/locations/%s", locationId)
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// WIP
func CreateLocation(network *net.Network, locationId string) (string, error) {
	url := fmt.Sprintf("/v1/locations", locationId)
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func LocationList(network *net.Network) ([]models.LocationSummary, error) {
	url := "/v1/locations"
	var locationList []models.LocationSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return locationList, err
	}

	err = json.Unmarshal(body, &locationList)
	return locationList, err
}
