package entity_effectors

import(
	"fmt"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func EffectorList(application, entity string) []models.EffectorSummary {
	url := "http://192.168.50.101:8081/v1/applications/" + application + "/entities/"+ entity + "/effectors"
	req := net.NewGetRequest(url)
	body, err := net.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	
	var effectorList []models.EffectorSummary
	err = json.Unmarshal(body, &effectorList)
	if err != nil{
		fmt.Println(err)
	}
	return effectorList
}