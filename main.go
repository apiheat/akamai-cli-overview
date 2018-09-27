package main

import (
	"log"
	"os"
	"sort"

	common "github.com/apiheat/akamai-cli-common"
	edgegrid "github.com/apiheat/go-edgegrid"

	"github.com/urfave/cli"
)

var (
	apiClient           *edgegrid.Client
	apiClientOpts       *edgegrid.ClientOptions
	appVer, appName     string
	groupID, contractID string
)

func main() {
	app := common.CreateNewApp(appName, "A CLI to interact with Akamai account information", appVer)
	app.Flags = common.CreateFlags()
	app.Before = func(c *cli.Context) error {

		apiClientOpts := &edgegrid.ClientOptions{}
		apiClientOpts.ConfigPath = c.GlobalString("config")
		apiClientOpts.ConfigSection = c.GlobalString("section")
		apiClientOpts.DebugLevel = c.GlobalString("debug")

		// create new Akamai API client
		apiClient = edgegrid.NewClient(nil, apiClientOpts)

		return nil
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
					Name:        "contractID",
					Usage:       "",
					Destination: &contractID,
				},
				cli.StringFlag{
					Name:        "groupID",
					Usage:       "",
					Destination: &groupID,
				},
				cli.BoolFlag{
					Name:  "all",
					Usage: "Show All properties",
				},
			},
			Action: cmdListProperties,
		},
		{
			Name:  "edge-hostnames",
			Usage: "Lists all edge hostnames available under a contract",
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
			Action: cmdListEdgeHostNames,
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
			Action: cmdListCPcodes,
		},
		{
			Name:   "rule-formats",
			Usage:  "List all available rule formats",
			Action: cmdListRules,
		},
		{
			Name:   "custom-overrides",
			Usage:  "Lists the set of custom XML metadata overrides configured for you by Akamai representatives",
			Action: cmdListOverrides,
		},
		{
			Name:   "custom-behaviors",
			Usage:  "Lists the set of custom XML metadata behaviors configured for you by Akamai representativess",
			Action: cmdListBehaviors,
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
