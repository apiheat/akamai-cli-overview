package main

import (
	"log"
	"os"
	"sort"

	edgegrid "github.com/RafPe/go-edgegrid"
	homedir "github.com/mitchellh/go-homedir"

	"github.com/urfave/cli"
)

var (
	apiClient                *edgegrid.Client
	apiClientOpts            *edgegrid.ClientOptions
	homeDir, output, version string
	groupID, contractID      string
)

// Constants
const (
	URL     = "/papi/v1"
	padding = 3
)

func main() {

	/*
		Sets default value for credentials configuration file
		to be pointing to ~/.edgerc
	*/
	homeDir, _ = homedir.Dir()
	homeDir += string(os.PathSeparator) + ".edgerc"

	/*
		Initialize values with using ENV variables either defaults
		AKAMAI_EDGERC_CONFIG  : for config file path
		AKAMAI_EDGERC_SECTION : for section
	*/
	apiClientOpts := &edgegrid.ClientOptions{}
	apiClientOpts.ConfigPath = getEnv(string(edgegrid.EnvVarEdgercPath), homeDir)
	apiClientOpts.ConfigSection = getEnv(string(edgegrid.EnvVarEdgercSection), "default")

	/*
		Sets default values for app and global flags
	*/
	appName := "akamai-overview"

	app := cli.NewApp()
	app.Name = appName
	app.HelpName = appName
	app.Usage = "A CLI to interact with Akamai account information"
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

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "section, s",
			Value:       "default",
			Usage:       "`NAME` of section to use from credentials file",
			Destination: &apiClientOpts.ConfigSection,
			EnvVar:      string(edgegrid.EnvVarEdgercSection),
		},
		cli.StringFlag{
			Name:        "config, c",
			Value:       homeDir,
			Usage:       "Location of the credentials `FILE`",
			Destination: &apiClientOpts.ConfigPath,
			EnvVar:      string(edgegrid.EnvVarEdgercPath),
		},
		cli.StringFlag{
			Name:        "output",
			Value:       "json",
			Usage:       "Defines output type ( json | table ) ",
			Destination: &output,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "contracts",
			Usage:  "List associated account contracts",
			Action: cmdListContracts,
		},
		{
			Name:   "groups",
			Usage:  "List associated account groups",
			Action: cmdListGroups,
		},
		{
			Name:  "products",
			Usage: "List associated contract products",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "contractID",
					Usage:       "",
					Destination: &contractID,
				},
			},
			Action: cmdListProducts,
		},
		{
			Name:  "properties",
			Usage: "Lists properties available for the current contract and group",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "groupID",
					Value: "",
					Usage: "Unique identifier for the group",
				},
				cli.StringFlag{
					Name:  "contractID",
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
			Name:   "products",
			Usage:  "Lists the set of products that are available under a given contract",
			Action: cmdProducts,
		},
		{
			Name:  "edge-hostnames",
			Usage: "Lists all edge hostnames available under a contract",
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
			Name:  "cpcodes",
			Usage: "List associated contract/group cpcodes",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "contractID",
					Usage:       "",
					Destination: &contractID,
				},
				cli.StringFlag{
					Name:        "groupID",
					Usage:       "",
					Destination: &groupID,
				},
			},
			Action:   cmdListCPcodes,
			Category: "CPCodes actions",
		},
		{
			Name:   "rule-formats",
			Usage:  "List all available rule formats",
			Action: cmdRules,
		},
		{
			Name:   "custom-overrides",
			Usage:  "Lists the set of custom XML metadata overrides configured for you by Akamai representatives",
			Action: cmdOverrides,
		},
		{
			Name:   "custom-behaviors",
			Usage:  "Lists the set of custom XML metadata behaviors configured for you by Akamai representativess",
			Action: cmdBehaviors,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Action = func(c *cli.Context) error {

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
