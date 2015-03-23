package catalog

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func Catalog(network *net.Network) []models.Application{
	url := "/v1/catalog/applications"
	req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	var applications []models.Application
	err = json.Unmarshal(body, &applications)
	return applications
}