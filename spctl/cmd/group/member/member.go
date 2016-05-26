package member

import "github.com/codegangsta/cli"

// Command holds the user subcommands
var Command = cli.Command{
	Name:        "member",
	Usage:       "manage groups memberships",
	ArgsUsage:   "",
	Description: "Manage SecurePass groups memberships.",
	Subcommands: []cli.Command{},
}
