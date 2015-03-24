package application

import(
	"fmt"
	"os"
	"path/filepath"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func Tree(network *net.Network) []models.Tree {
	url := "/v1/applications/tree"
	req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	
	var tree []models.Tree
	err = json.Unmarshal(body, &tree)
	if err != nil{
		fmt.Println(err)
	}
	return tree
}

func Application(network *net.Network, app string)  models.ApplicationSummary {
	url := fmt.Sprintf("/v1/applications/%s", app)
	req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	
	var appSummary models.ApplicationSummary
	err = json.Unmarshal(body, &appSummary)
	if err != nil{
		fmt.Println(err)
	}
	return appSummary
}

func Applications(network *net.Network)  []models.ApplicationSummary {
	url := fmt.Sprintf("/v1/applications")
	req := network.NewGetRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	
	var appSummary []models.ApplicationSummary
	err = json.Unmarshal(body, &appSummary)
	if err != nil{
		fmt.Println(err)
	}
	return appSummary
}

func Create(network *net.Network, filePath string) models.TaskSummary{
	url := "/v1/applications"
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
	var response models.TaskSummary
	err = json.Unmarshal(body, &response)
	if err != nil{
		fmt.Println(err)
	}
	return response
}

func Delete(network *net.Network, application string) models.TaskSummary{
	url := fmt.Sprintf("/v1/applications/%s", application)
	req := network.NewDeleteRequest(url)
	body, err := network.SendRequest(req)
	if err != nil {
		fmt.Println(err)
	}
	var response models.TaskSummary
	err = json.Unmarshal(body, &response)
	if err != nil{
		fmt.Println(err)
	}
	return response
}