package locations

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func LocationList() []models.LocationSummary {
	url := "http://192.168.50.101:8081/v1/locations"
	req := net.NewGetRequest(url)
	body, err := net.SendRequest(req)
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