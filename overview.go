package main

import (
	"fmt"

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
	contracts, _, err := apiClient.PropertyAPI.ListPropertyAPIContracts()
	if err != nil {
		fmt.Println(err)
	}

	OutputJSON(contracts)

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
	groups, _, err := apiClient.PropertyAPI.ListPropertyAPIGroups()
	if err != nil {
		fmt.Println(err)
	}

	OutputJSON(groups)

	return nil
}

/*
	listProducts
*/
func cmdListProducts(c *cli.Context) error {
	return listProducts(c)
}

func listProducts(c *cli.Context) error {
	verifyArgumentByName(c, "contractID")

	products, _, err := apiClient.PropertyAPI.ListPropertyAPIProducts(contractID)
	if err != nil {
		fmt.Println(err)
	}

	OutputJSON(products)

	return nil
}

/*
	listCPcodes
*/
func cmdListCPcodes(c *cli.Context) error {
	return listCPcodes(c)
}

func listCPcodes(c *cli.Context) error {
	verifyArgumentByName(c, "contractID")
	verifyArgumentByName(c, "groupID")

	cpcodes, _, err := apiClient.PropertyAPI.ListPropertyAPICPCodes(contractID, groupID)
	if err != nil {
		fmt.Println(err)
	}

	OutputJSON(cpcodes)

	return nil
}

/*
	ListEdgeHostNames
*/
func cmdListEdgeHostNames(c *cli.Context) error {
	return listEdgeHostNames(c)
}

func listEdgeHostNames(c *cli.Context) error {
	verifyArgumentByName(c, "contractID")
	verifyArgumentByName(c, "groupID")

	edgeHosts, _, err := apiClient.PropertyAPI.ListPropertyAPICPEdgehosts(contractID, groupID)
	if err != nil {
		fmt.Println(err)
	}

	OutputJSON(edgeHosts)

	return nil
}

/*
	ListProperties
*/
func cmdListProperties(c *cli.Context) error {
	return listProperties(c)
}

func listProperties(c *cli.Context) error {
	verifyArgumentByName(c, "contractID")
	verifyArgumentByName(c, "groupID")

	allProperties, _, err := apiClient.PropertyAPI.ListPropertyAPIProperties(contractID, groupID)
	if err != nil {
		fmt.Println(err)
	}

	OutputJSON(allProperties)

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

	OutputJSON(k.RuleFormats.Items)

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

	OutputJSON(k)

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

	OutputJSON(k)

	return nil
}
