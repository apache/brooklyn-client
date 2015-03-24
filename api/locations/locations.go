package locations

import (
	"encoding/json"
	"fmt"
	"github.com/robertgmoss/brooklyn-cli/models"
	"github.com/robertgmoss/brooklyn-cli/net"
)

func LocationList(network *net.Network) []models.LocationSummary {
	url := "/v1/locations"
	req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
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
