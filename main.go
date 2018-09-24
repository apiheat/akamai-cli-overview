package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	edgegrid "github.com/RafPe/go-edgegrid"
	homedir "github.com/mitchellh/go-homedir"

	"github.com/urfave/cli"
)

var (
	apiClient           *edgegrid.Client
	apiClientOpts       *edgegrid.ClientOptions
	homeDir, appVer     string
	groupID, contractID string
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
	app.Version = appVer
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
	}

	app.Before = func(c *cli.Context) error {

		// create new Akamai API client
		apiClient = edgegrid.NewClient(nil, apiClientOpts)

		// if err != nil {
		// 	return cli.NewExitError(errorProfile, 1)
		// }

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
			Action:   cmdListCPcodes,
			Category: "CPCodes actions",
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

func verifyArgumentByName(c *cli.Context, argName string) {
	if c.String(argName) == "" {
		log.Fatal(fmt.Sprintf("Please provide required argument(s)! [ %s ]", argName))
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
