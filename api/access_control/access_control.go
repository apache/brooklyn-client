package access_control

import(
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func Access(network *net.Network) models.AccessSummary {
	url := fmt.Sprintf("/v1/access")
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	
	var access models.AccessSummary
	err = json.Unmarshal(body, &access)
	if err != nil {
		fmt.Println(err)
	}
	return access
}