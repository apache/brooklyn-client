package locations

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func LocationList(network *net.Network,) []models.LocationSummary {
	url := "/v1/locations"
	req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	
	var locationList []models.LocationSummary
	err = json.Unmarshal(body, &locationList)
	if err != nil{
		fmt.Println(err)
	}
	return locationList
}