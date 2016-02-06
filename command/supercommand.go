package command

// A command with further (sub) commands, like 'git remote', with its 'git remote add' etc.
type SuperCommand interface {
	Command

	// Get the sub command wih the given name
	SubCommand(name string) Command

	// Get the names of all subcommands
	SubCommandNames() []string
}
