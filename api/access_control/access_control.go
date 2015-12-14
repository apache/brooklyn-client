package access_control

import(
	"encoding/json"
	"fmt"
	"github.com/brooklyncentral/brooklyn-cli/models"
	"github.com/brooklyncentral/brooklyn-cli/net"
)

func Access(network *net.Network) (models.AccessSummary, error) {
	url := fmt.Sprintf("/v1/access")
    var access models.AccessSummary

    body, err := network.SendGetRequest(url)
    if err != nil {
        return access, err
    }

	err = json.Unmarshal(body, &access)
	if err != nil {
		return access, err
	}
	return access, nil
}

// WIP
//func LocationProvisioningAllowed(network *net.Network, allowed bool) {
//	url := fmt.Sprintf("/v1/access/locationProvisioningAllowed")
//	body, err := network.SendPostRequest(url)
//	if err != nil {
//		fmt.Println(err)
//	}
//}