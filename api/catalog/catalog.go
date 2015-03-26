package catalog

import (
	"encoding/json"
	"fmt"
	"github.com/robertgmoss/brooklyn-cli/models"
	"github.com/robertgmoss/brooklyn-cli/net"
	"os"
	"path/filepath"
)

func Catalog(network *net.Network) []models.Application {
	url := "/v1/catalog/applications"
	body, err := network.SendGetRequest(url)
	if err != nil {
		fmt.Println(err)
	}
	var applications []models.Application
	err = json.Unmarshal(body, &applications)
	return applications
}

func AddCatalog(network *net.Network, filePath string) string {
	url := "/v1/catalog"
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	req := network.NewPostRequest(url, file)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}
