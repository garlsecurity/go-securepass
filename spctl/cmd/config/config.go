package config

import (
	"io"
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/garlsecurity/securepassctl"
	"github.com/garlsecurity/securepassctl/spctl/service"
)

// Command holds the config command
var Command = cli.Command{
	Name:        "config",
	Usage:       "configure SecurePass",
	ArgsUsage:   " ",
	Description: "Create or update SecurePass configuration.",
	Action:      ActionConfig,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Configuration file",
			Value: "/etc/securepass.conf",
		},
		cli.StringFlag{
			Name:  "appid, i",
			Usage: "Application ID",
		},
		cli.StringFlag{
			Name:  "endpoint, e",
			Usage: "Endpoint URL",
			Value: securepassctl.DefaultRemote,
		},
		cli.StringFlag{
			Name:  "appsecret, s",
			Usage: "App secret",
		},
		cli.StringFlag{
			Name:  "realm, r",
			Usage: "Default Domain/Realm (and allow NSS login)",
		},
		cli.StringFlag{
			Name:  "root",
			Usage: "Coma separated list of allowed root users",
		},
		cli.BoolFlag{
			Name:  "stdout",
			Usage: "Write to stdout instead of file",
		},
	},
}

// ActionConfig is the config command handler
func ActionConfig(c *cli.Context) {
	var writer io.Writer = os.Stdout

	if !c.Bool("stdout") {
		w, err := os.OpenFile(c.String("config"),
			os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		writer = w
	}

	s := service.Service
	nss := service.NSSSettings
	ssh := service.SSHSettings
	if c.String("appid") != "" {
		s.AppID = c.String("appid")
	}
	if c.String("endpoint") != "" {
		s.Endpoint = c.String("endpoint")
	}
	if c.String("appsecret") != "" {
		s.AppSecret = c.String("appsecret")
	}
	if c.String("realm") != "" {
		nss.Realm = c.String("realm")
	}
	if c.String("root") != "" {
		ssh.Root = c.String("root")
	}

	service.WriteConfiguration(writer, s, nss, ssh)
}
