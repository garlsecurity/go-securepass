package realm

import "github.com/urfave/cli"

// Command holds the realm subcommands
var Command = cli.Command{
	Name:        "realm",
	Usage:       "manage realm settings",
	ArgsUsage:   "",
	Description: "Manage realm settings.",
	Subcommands: []cli.Command{},
}
