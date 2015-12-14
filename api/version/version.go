package version

import (
	"github.com/brooklyncentral/brooklyn-cli/net"
    "github.com/brooklyncentral/brooklyn-cli/models"
    "encoding/json"
)

func Version(network *net.Network) (models.VersionSummary, error) {
	url := "/v1/server/version"
    var versionSummary models.VersionSummary
	body, err := network.SendGetRequest(url)
	if err != nil {
		return versionSummary, err
	}
    err = json.Unmarshal(body, &versionSummary)
	return versionSummary, err
}
