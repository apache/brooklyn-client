package access_control

import (
	"encoding/json"
	"fmt"
	"github.com/apache/brooklyn-client/models"
	"github.com/apache/brooklyn-client/net"
)

func Access(network *net.Network) (models.AccessSummary, error) {
	url := fmt.Sprintf("/v1/access")
	var access models.AccessSummary

	body, err := network.SendGetRequest(url)
	if err != nil {
		return access, err
	}

	err = json.Unmarshal(body, &access)
	return access, err
}

// WIP
//func LocationProvisioningAllowed(network *net.Network, allowed bool) {
//	url := fmt.Sprintf("/v1/access/locationProvisioningAllowed")
//	body, err := network.SendPostRequest(url)
//	if err != nil {
//		error_handler.ErrorExit(err)
//	}
//}
