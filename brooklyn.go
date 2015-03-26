package main

import (
	"github.com/robertgmoss/brooklyn-cli/app"
	"github.com/robertgmoss/brooklyn-cli/command_factory"
	"github.com/robertgmoss/brooklyn-cli/command_runner"
	"github.com/robertgmoss/brooklyn-cli/net"
	"github.com/robertgmoss/brooklyn-cli/io"
	"os"
	"path/filepath"
)

func getNetworkCredentialsFromConfig(yamlMap map[string]interface{}) (string, string, string){
	var target, username, password string
	target, found := yamlMap["target"].(string)
	if found {
		auth, found := yamlMap["auth"].(map[string]interface{})
		if found {
			creds := auth[target].(map[string]interface{})
			username, found = creds["username"].(string)
			if found {
				password, found = creds["password"].(string)
			}
		}
	}
	return target, username, password
}

func main() {
	config := io.GetConfig()
	target, username, password := getNetworkCredentialsFromConfig(config.Map) 
	//target, username, password := "http://192.168.50.101:8081", "brooklyn", "Sns4Hh9j7l"
	network := net.NewNetwork(target, username, password)
	cmdFactory := command_factory.NewFactory(network, config)
	cmdRunner := command_runner.NewRunner(cmdFactory)
	metaDatas := cmdFactory.CommandMetadatas()
	theApp := app.NewApp(filepath.Base(os.Args[0]), cmdRunner, metaDatas...)
	theApp.Run(os.Args)
}
