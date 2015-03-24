package catalog

import(
	"fmt"
	"os"
	"path/filepath"
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

func AddCatalog(network *net.Network, filePath string) string{
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