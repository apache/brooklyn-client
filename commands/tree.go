package commands

import(
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/robertgmoss/brooklyn-cli/api/application"
	"github.com/robertgmoss/brooklyn-cli/command_metadata"
	"github.com/robertgmoss/brooklyn-cli/models"
	"github.com/robertgmoss/brooklyn-cli/net"
)

type Tree struct {
	network *net.Network
}

func NewTree(network *net.Network) (cmd *Tree){
	cmd = new(Tree)
	cmd.network = network
	return
}

func (cmd *Tree) Metadata() command_metadata.CommandMetadata {
	return command_metadata.CommandMetadata{
		Name:        "tree",
		Description: "show the tree of all applications",
		Usage:       "BROOKLYN_NAME tree",
		Flags: []cli.Flag{},
	}
}	

func (cmd *Tree) Run(c *cli.Context) {
	trees := application.Tree(cmd.network)
	cmd.printTrees(trees, "")
}

func (cmd *Tree) printTrees(trees []models.Tree, indent string){
	for i, app := range trees {
		cmd.printTree(app, indent, i == len(trees) - 1)
	}
}

func (cmd *Tree) printTree(tree models.Tree, indent string, last bool){
	fmt.Println(indent + "|-", tree.Name)
	fmt.Println(indent+ "+-", tree.Type)
	
	if last {
		indent = indent + "  "
	}else {
		indent = indent + "| "
	}
	cmd.printTrees(tree.Children, indent)
}