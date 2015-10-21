package locations

import (
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func LocatedLocations(network *net.Network) string {
	url := "/v1/locations/usage/LocatedLocations"
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func GetLocation(network *net.Network, locationId string) string {
	url := fmt.Sprintf("/v1/locations/%s", locationId)
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func DeleteLocation(network *net.Network, locationId string) string {
	url := fmt.Sprintf("/v1/locations/%s", locationId)
	body, err := network.SendDeleteRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

// WIP
func CreateLocation(network *net.Network, locationId string) string {
	url := fmt.Sprintf("/v1/locations", locationId)
	body, err := network.SendEmptyPostRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func LocationList(network *net.Network) []models.LocationSummary {
	url := "/v1/locations"
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	var locationList []models.LocationSummary
	err = json.Unmarshal(body, &locationList)
	if err != nil {
		fmt.Println(err)
	}
	return locationList
}
