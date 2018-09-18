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

	if output == "json" {
		OutputJSON(contracts)
	} else {
		outputTableContracts(contracts)
	}

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

	if output == "json" {
		OutputJSON(groups)
	} else {
		outputTableGroups(groups)
	}

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

	if output == "json" {
		OutputJSON(products)
	} else {
		outputTableProducts(products)
	}

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

	if output == "json" {
		OutputJSON(cpcodes)
	} else {
		outputTableCPCodes(cpcodes)
	}

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

	cpcodes, _, err := apiClient.PropertyAPI.ListPropertyAPICPEdgehosts(contractID, groupID)
	if err != nil {
		fmt.Println(err)
	}

	if output == "json" {
		OutputJSON(cpcodes)
	} else {
		// outputTableCPCodes(cpcodes)
	}

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

	cpcodes, _, err := apiClient.PropertyAPI.ListPropertyAPIProperties(contractID, groupID)
	if err != nil {
		fmt.Println(err)
	}

	if output == "json" {
		OutputJSON(cpcodes)
	} else {
		// outputTableCPCodes(cpcodes)
	}

	return nil
}
