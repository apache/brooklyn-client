package main

import (
	"github.com/brooklyncentral/brooklyn-cli/app"
	"github.com/brooklyncentral/brooklyn-cli/command_factory"
	"github.com/brooklyncentral/brooklyn-cli/command_runner"
	"github.com/brooklyncentral/brooklyn-cli/net"
	"github.com/brooklyncentral/brooklyn-cli/io"
	"os"
	"path/filepath"
	"github.com/brooklyncentral/brooklyn-cli/scope"
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

	args, scope := scope.ScopeArguments(os.Args)
	cmdRunner := command_runner.NewRunner(scope, cmdFactory)
	metaDatas := cmdFactory.CommandMetadatas()
	theApp := app.NewApp(filepath.Base(args[0]), cmdRunner, metaDatas...)
	theApp.Run(args)
}
