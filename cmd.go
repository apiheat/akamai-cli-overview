package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	client "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
	"github.com/urfave/cli"
)

func cmdGroups(c *cli.Context) error {
	return listGroups(c)
}

func cmdRules(c *cli.Context) error {
	return listRules(c)
}

func cmdOverrides(c *cli.Context) error {
	return listOverrides(c)
}

func cmdCPCodes(c *cli.Context) error {
	if c.String("group") == "" {
		log.Fatal("Please provide Group ID")
	}
	if c.String("contract") == "" {
		log.Fatal("Please provide Contract ID")
	}

	return listCPCodes(c)
}

func cmdEdges(c *cli.Context) error {
	if c.String("group") == "" {
		log.Fatal("Please provide Group ID")
	}
	if c.String("contract") == "" {
		log.Fatal("Please provide Contract ID")
	}

	return listEdges(c)
}

func cmdContracts(c *cli.Context) error {
	return listContracts(c)
}

func cmdProducts(c *cli.Context) error {
	return listProducts(c)
}

func cmdProperties(c *cli.Context) error {
	if c.String("group") == "" {
		log.Fatal("Please provide Group ID")
	}
	if c.String("contract") == "" {
		log.Fatal("Please provide Contract ID")
	}

	return listProperties(c)
}

func listEdges(c *cli.Context) error {
	urlStr := fmt.Sprintf("%s/edgehostnames?contractId=%s&groupId=%s&options=%s", URL, c.String("contract"), c.String("group"), c.String("options"))

	if debug {
		println(urlStr)
	}

	data, _ := fetchData(urlStr, "GET", nil)

	if debug {
		println(data)
	}

	result, err := edgesRespParse(data)
	errorCheck(err)

	outputTableEdges(result)

	return nil
}

func listCPCodes(c *cli.Context) error {
	urlStr := fmt.Sprintf("%s/cpcodes?contractId=%s&groupId=%s", URL, c.String("contract"), c.String("group"))

	if debug {
		println(urlStr)
	}

	data, _ := fetchData(urlStr, "GET", nil)

	if debug {
		println(data)
	}

	result, err := cpcodesRespParse(data)
	errorCheck(err)

	outputTableCPCodes(result)

	return nil
}

func listProperties(c *cli.Context) error {
	if c.Bool("all") {
		fmt.Println("To be implemented for all properties in all groups for all contracts")
	} else {
		urlStr := fmt.Sprintf("%s/properties?contractId=%s&groupId=%s", URL, c.String("contract"), c.String("group"))

		if debug {
			println(urlStr)
		}

		data, _ := fetchData(urlStr, "GET", nil)

		if debug {
			println(data)
		}

		result, err := propertiesRespParse(data)
		errorCheck(err)

		outputTableProperties(result)
	}

	return nil
}

func listProducts(c *cli.Context) error {
	id, err := setID(c)
	if err != nil {
		log.Fatal("Please provide contract ID")
	}
	urlStr := fmt.Sprintf("%s/products?contractId=%s", URL, id)

	if debug {
		println(urlStr)
	}

	data, _ := fetchData(urlStr, "GET", nil)

	if debug {
		println(data)
	}

	result, err := productsRespParse(data)
	errorCheck(err)

	outputTableProducts(result)

	return nil
}

func listContracts(c *cli.Context) error {
	urlStr := fmt.Sprintf("%s/contracts", URL)

	if debug {
		println(urlStr)
	}

	data, _ := fetchData(urlStr, "GET", nil)

	if debug {
		println(data)
	}

	result, err := contractsRespParse(data)
	errorCheck(err)

	outputTableContracts(result)

	return nil
}

func listGroups(c *cli.Context) error {
	urlStr := fmt.Sprintf("%s/groups", URL)

	if debug {
		println(urlStr)
	}

	data, _ := fetchData(urlStr, "GET", nil)

	if debug {
		println(data)
	}

	result, err := groupsRespParse(data)
	errorCheck(err)

	outputTableGroups(result)

	return nil
}

func listRules(c *cli.Context) error {
	urlStr := fmt.Sprintf("%s/rule-formats", URL)

	if debug {
		println(urlStr)
	}

	data, _ := fetchData(urlStr, "GET", nil)

	if debug {
		println(data)
	}

	result, err := rulesRespParse(data)
	errorCheck(err)

	outputTableRules(result)

	return nil
}

func listOverrides(c *cli.Context) error {
	urlStr := fmt.Sprintf("%s/custom-overrides", URL)

	if debug {
		println(urlStr)
	}

	data, _ := fetchData(urlStr, "GET", nil)

	if debug {
		println(data)
	}

	result, err := overridesRespParse(data)
	errorCheck(err)

	outputTableOverrides(result)

	return nil
}

func rulesRespParse(in string) (data RulesResponse, err error) {
	if err = json.Unmarshal([]byte(in), &data); err != nil {
		return
	}
	return
}

func overridesRespParse(in string) (data OverridesResponse, err error) {
	if err = json.Unmarshal([]byte(in), &data); err != nil {
		return
	}
	return
}

func cpcodesRespParse(in string) (data CPCodesResponse, err error) {
	if err = json.Unmarshal([]byte(in), &data); err != nil {
		return
	}
	return
}

func groupsRespParse(in string) (data GroupsResponse, err error) {
	if err = json.Unmarshal([]byte(in), &data); err != nil {
		return
	}
	return
}

func contractsRespParse(in string) (data ContractsResponse, err error) {
	if err = json.Unmarshal([]byte(in), &data); err != nil {
		return
	}
	return
}

func productsRespParse(in string) (data ProductsResponse, err error) {
	if err = json.Unmarshal([]byte(in), &data); err != nil {
		return
	}
	return
}

func propertiesRespParse(in string) (data PropertiesResponse, err error) {
	if err = json.Unmarshal([]byte(in), &data); err != nil {
		return
	}
	return
}

func edgesRespParse(in string) (data EdgesResponse, err error) {
	if err = json.Unmarshal([]byte(in), &data); err != nil {
		return
	}
	return
}

func errorCheck(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func printJSON(str interface{}) {
	jsonRes, _ := json.MarshalIndent(str, "", "  ")
	fmt.Printf("%+v\n", string(jsonRes))
}

func fetchData(urlPath, method string, body io.Reader) (string, int) {
	req, err := client.NewRequest(edgeConfig, method, urlPath, body)
	errorCheck(err)

	resp, err := client.Do(edgeConfig, req)
	errorCheck(err)

	defer resp.Body.Close()
	byt, _ := ioutil.ReadAll(resp.Body)

	return string(byt), resp.StatusCode
}

func setID(c *cli.Context) (string, error) {
	var id string
	if c.NArg() == 0 {
		return "", errors.New("ID not present")
	}

	id = c.Args().Get(0)
	return id, nil
}
