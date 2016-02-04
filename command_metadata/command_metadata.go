package command_metadata

import "github.com/codegangsta/cli"

type CommandMetadata struct {
	Name            string
	Aliases         []string
	ShortName       string
	Usage           string
	Description     string
	Flags           []cli.Flag
	SkipFlagParsing bool
	TotalArgs       int //Optional: number of required arguments to skip for flag verification
	Operands        []CommandMetadata
}
