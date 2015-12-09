# Brooklyn CLI

## Compiling

1. Ensure your [$GOPATH](http://golang.org/cmd/go/#hdr-GOPATH_environment_variable) is set correctly 
2. Get the cli source code: go get github.com/brooklyncentral/brooklyn-cli
3. cd $GOPATH/src/github.com/brooklyncentral/brooklyn-cli/br
4. go build

## Running
you will need to login to your Brooklyn instance with
```
  $ ./brooklyn-cli login URL USER PASSWORD
```  
See the help command for info on all commands
```
  $ ./brooklyn-cli help
```  
Or for help on individual commands
```
  $ ./brooklyn-cli help COMMAND
```
