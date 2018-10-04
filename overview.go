package main

import (
	"fmt"

	common "github.com/apiheat/akamai-cli-common"
	"github.com/urfave/cli"
)

/*
	listContracts
*/
func cmdListContracts(c *cli.Context) error {
	return listContracts(c)
}

func listContracts(c *cli.Context) error {
	// List all network lists
	contracts, _, err := apiClient.Property.ListPropertyContracts()
	if err != nil {
		fmt.Println(err)
	}

	common.OutputJSON(contracts)

	return nil
}

/*
	listGroups
*/
func cmdListGroups(c *cli.Context) error {
	return listGroups(c)
}

func listGroups(c *cli.Context) error {
	// List all network lists
	groups, _, err := apiClient.Property.ListPropertyGroups()
	if err != nil {
		fmt.Println(err)
	}

	common.OutputJSON(groups)

	return nil
}

/*
	listProducts
*/
func cmdListProducts(c *cli.Context) error {
	return listProducts(c)
}

func listProducts(c *cli.Context) error {
	common.VerifyArgumentByName(c, "contractID")

	products, _, err := apiClient.Property.ListPropertyProducts(contractID)
	if err != nil {
		fmt.Println(err)
	}

	common.OutputJSON(products)

	return nil
}

/*
	listCPcodes
*/
func cmdListCPcodes(c *cli.Context) error {
	return listCPcodes(c)
}

func listCPcodes(c *cli.Context) error {
	common.VerifyArgumentByName(c, "contractID")
	common.VerifyArgumentByName(c, "groupID")

	cpcodes, _, err := apiClient.Property.ListPropertyCPCodes(contractID, groupID)
	if err != nil {
		fmt.Println(err)
	}

	common.OutputJSON(cpcodes)

	return nil
}

/*
	ListEdgeHostNames
*/
func cmdListEdgeHostNames(c *cli.Context) error {
	return listEdgeHostNames(c)
}

func listEdgeHostNames(c *cli.Context) error {
	common.VerifyArgumentByName(c, "contractID")
	common.VerifyArgumentByName(c, "groupID")

	edgeHosts, _, err := apiClient.Property.ListPropertyCPEdgehosts(contractID, groupID)
	if err != nil {
		fmt.Println(err)
	}

	common.OutputJSON(edgeHosts)

	return nil
}

/*
	ListProperties
*/
func cmdListProperties(c *cli.Context) error {
	return listProperties(c)
}

func listProperties(c *cli.Context) error {
	common.VerifyArgumentByName(c, "contractID")
	common.VerifyArgumentByName(c, "groupID")

	allProperties, _, err := apiClient.Property.ListPropertyProperties(contractID, groupID)
	if err != nil {
		fmt.Println(err)
	}

	common.OutputJSON(allProperties)

	return nil
}

/*
	#TODO:The below functions still should end up in go-edge client in the papi secion ...
*/

/*
	ListRules
*/
func cmdListRules(c *cli.Context) error {
	return listRules(c)
}

func listRules(c *cli.Context) error {
	var k *RulesResponse
	_, err := apiClient.NewRequest("GET", "/papi/v1/rule-formats", nil, &k)
	if err != nil {

		return err
	}

	common.OutputJSON(k.RuleFormats.Items)

	return nil
}

/*
	ListOverrides
*/
func cmdListOverrides(c *cli.Context) error {
	return listOverrides(c)
}

func listOverrides(c *cli.Context) error {
	var k *OverridesResponse
	_, err := apiClient.NewRequest("GET", "/papi/v1/custom-overrides", nil, &k)
	if err != nil {

		return err
	}

	common.OutputJSON(k)

	return nil
}

/*
	ListBehaviors
*/
func cmdListBehaviors(c *cli.Context) error {
	return listBehaviors(c)
}

func listBehaviors(c *cli.Context) error {
	var k *OverridesResponse
	_, err := apiClient.NewRequest("GET", "/papi/v1/custom-behaviors", nil, &k)
	if err != nil {

		return err
	}

	common.OutputJSON(k)

	return nil
}
