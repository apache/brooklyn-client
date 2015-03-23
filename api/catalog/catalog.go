package catalog

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func Catalog() []models.Application{
	url := "http://192.168.50.101:8081/v1/catalog/applications"
	req := net.NewGetRequest(url)
	body, err := net.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	var applications []models.Application //[]map[string]interface{}
	err = json.Unmarshal(body, &applications)
	return applications
}