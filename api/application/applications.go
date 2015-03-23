package application

import(
	"fmt"
	"os"
	"path/filepath"
	"encoding/json"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/models"
)

func Tree() []models.Tree {
	url := "http://192.168.50.101:8081/v1/applications/tree"
	req := net.NewGetRequest(url)
	body, err := net.SendRequest(req)
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

func Application(app string)  models.ApplicationSummary {
	url := "http://192.168.50.101:8081/v1/applications/" + app
	req := net.NewGetRequest(url)
	body, err := net.SendRequest(req)
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

func Create(filePath string) models.TaskSummary{
	url := "http://192.168.50.101:8081/v1/applications"
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	req := net.NewPostRequest(url, file)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	body, err := net.SendRequest(req)
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