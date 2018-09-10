package main

import (
	"os"
	"sort"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

var (
	colorOn, raw, debug       bool
	version, appName          string
	configSection, configFile string
	edgeConfig                edgegrid.Config
)

// Constants
const (
	URL     = "/papi/v1"
	padding = 3
)

func main() {
	_, inCLI := os.LookupEnv("AKAMAI_CLI")

	appName = "akamai-overview"
	if inCLI {
		appName = "akamai overview"
	}

	app := cli.NewApp()
	app.Name = appName
	app.HelpName = appName
	app.Usage = "A CLI to list contracts, groups, properties"
	app.Version = version
	app.Copyright = ""
	app.Authors = []cli.Author{
		{
			Name: "Petr Artamonov",
		},
		{
			Name: "Rafal Pieniazek",
		},
	}

	dir, _ := homedir.Dir()
	dir += string(os.PathSeparator) + ".edgerc"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "section, s",
			Value:       "default",
			Usage:       "`NAME` of section to use from credentials file",
			Destination: &configSection,
			EnvVar:      "AKAMAI_EDGERC_SECTION",
		},
		cli.StringFlag{
			Name:        "config, c",
			Value:       dir,
			Usage:       "Location of the credentials `FILE`",
			Destination: &configFile,
			EnvVar:      "AKAMAI_EDGERC",
		},
		cli.BoolFlag{
			Name:        "no-color",
			Usage:       "Disable color output",
			Destination: &colorOn,
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Debug info",
			Destination: &debug,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "groups",
			Aliases: []string{"grp"},
			Usage:   "Lists the set of groups for account",
			Action:  cmdGroups,
		},
		{
			Name:    "contracts",
			Aliases: []string{"ctr"},
			Usage:   "Lists the set of contracts",
			Action:  cmdContracts,
		},
		{
			Name:    "properties",
			Aliases: []string{"prp"},
			Usage:   "Lists properties available for the current contract and group",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "group",
					Value: "",
					Usage: "Unique identifier for the group",
				},
				cli.StringFlag{
					Name:  "contract",
					Value: "",
					Usage: "Unique identifier for the contract",
				},
				cli.BoolFlag{
					Name:  "all",
					Usage: "Show All properties",
				},
			},
			Action: cmdProperties,
		},
		{
			Name:    "products",
			Aliases: []string{"pr"},
			Usage:   "Lists the set of products that are available under a given contract",
			Action:  cmdProducts,
		},
		{
			Name:    "edge-hostnames",
			Aliases: []string{"e"},
			Usage:   "Lists all edge hostnames available under a contract",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "group",
					Value: "",
					Usage: "Unique identifier for the group",
				},
				cli.StringFlag{
					Name:  "contract",
					Value: "",
					Usage: "Unique identifier for the contract",
				},
				cli.StringFlag{
					Name:  "options",
					Value: "mapDetails",
					Usage: "Comma-separated list of options to enable; mapDetails enables extra mapping-related information",
				},
			},
			Action: cmdEdges,
		},
		{
			Name:    "cpcodes",
			Aliases: []string{"cpc"},
			Usage:   "Lists CP codes available within your contract/group pairing",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "group",
					Value: "",
					Usage: "Unique identifier for the group",
				},
				cli.StringFlag{
					Name:  "contract",
					Value: "",
					Usage: "Unique identifier for the contract",
				},
			},
			Action: cmdCPCodes,
		},
		{
			Name:    "rule-formats",
			Aliases: []string{"rf"},
			Usage:   "List all available rule formats",
			Action:  cmdRules,
		},
		{
			Name:   "custom-overrides",
			Usage:  "Lists the set of custom XML metadata overrides configured for you by Akamai representatives",
			Action: cmdOverrides,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Before = func(c *cli.Context) error {
		if c.Bool("no-color") {
			color.NoColor = true
		}

		config(configFile, configSection)
		return nil
	}

	app.Run(os.Args)
}
