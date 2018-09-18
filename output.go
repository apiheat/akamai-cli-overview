package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	edgegrid "github.com/RafPe/go-edgegrid"
)

type GroupsResponse struct {
	AccountID   string `json:"accountId"`
	AccountName string `json:"accountName"`
	Groups      struct {
		Items []struct {
			GroupName     string   `json:"groupName"`
			GroupID       string   `json:"groupId"`
			ContractIds   []string `json:"contractIds,omitempty"`
			ParentGroupID string   `json:"parentGroupId,omitempty"`
		} `json:"items"`
	} `json:"groups"`
}

type ContractsResponse struct {
	AccountID string `json:"accountId"`
	Contracts struct {
		Items []struct {
			ContractID       string `json:"contractId"`
			ContractTypeName string `json:"contractTypeName"`
		} `json:"items"`
	} `json:"contracts"`
}

type ProductsResponse struct {
	AccountID  string `json:"accountId"`
	ContractID string `json:"contractId"`
	Products   struct {
		Items []struct {
			ProductName string `json:"productName"`
			ProductID   string `json:"productId"`
		} `json:"items"`
	} `json:"products"`
}

type PropertiesResponse struct {
	Properties struct {
		Items []struct {
			AccountID         string `json:"accountId"`
			ContractID        string `json:"contractId"`
			GroupID           string `json:"groupId"`
			PropertyID        string `json:"propertyId"`
			PropertyName      string `json:"propertyName"`
			LatestVersion     int    `json:"latestVersion"`
			StagingVersion    int    `json:"stagingVersion"`
			ProductionVersion int    `json:"productionVersion"`
			AssetID           string `json:"assetId"`
			Note              string `json:"note"`
		} `json:"items"`
	} `json:"properties"`
}

type EdgesResponse struct {
	AccountID     string `json:"accountId"`
	ContractID    string `json:"contractId"`
	GroupID       string `json:"groupId"`
	EdgeHostnames struct {
		Items []struct {
			EdgeHostnameID         string `json:"edgeHostnameId"`
			EdgeHostnameDomain     string `json:"edgeHostnameDomain"`
			ProductID              string `json:"productId"`
			DomainPrefix           string `json:"domainPrefix"`
			DomainSuffix           string `json:"domainSuffix"`
			Status                 string `json:"status,omitempty"`
			Secure                 bool   `json:"secure"`
			IPVersionBehavior      string `json:"ipVersionBehavior"`
			MapDetailsSerialNumber int    `json:"mapDetails:serialNumber"`
			MapDetailsSlotNumber   int    `json:"mapDetails:slotNumber,omitempty"`
			MapDetailsMapDomain    string `json:"mapDetails:mapDomain"`
		} `json:"items"`
	} `json:"edgeHostnames"`
}

type CPCodesResponse struct {
	AccountID  string `json:"accountId"`
	ContractID string `json:"contractId"`
	GroupID    string `json:"groupId"`
	Cpcodes    struct {
		Items []struct {
			CpcodeID    string    `json:"cpcodeId"`
			CpcodeName  string    `json:"cpcodeName"`
			ProductIds  []string  `json:"productIds"`
			CreatedDate time.Time `json:"createdDate"`
		} `json:"items"`
	} `json:"cpcodes"`
}

type RulesResponse struct {
	RuleFormats struct {
		Items []string `json:"items"`
	} `json:"ruleFormats"`
}

type OverridesResponse struct {
	AccountID       string `json:"accountId"`
	CustomOverrides struct {
		Items []struct {
			OverrideID    string    `json:"overrideId"`
			DisplayName   string    `json:"displayName"`
			Description   string    `json:"description"`
			Name          string    `json:"name"`
			Status        string    `json:"status"`
			UpdatedByUser string    `json:"updatedByUser"`
			UpdatedDate   time.Time `json:"updatedDate"`
		} `json:"items"`
	} `json:"customOverrides"`
}

type BehaviorsResponse struct {
	AccountID       string `json:"accountId"`
	CustomBehaviors struct {
		Items []struct {
			BehaviorID    string    `json:"behaviorId"`
			Name          string    `json:"name"`
			Status        string    `json:"status"`
			DisplayName   string    `json:"displayName"`
			Description   string    `json:"description"`
			UpdatedDate   time.Time `json:"updatedDate"`
			UpdatedByUser string    `json:"updatedByUser"`
		} `json:"items"`
	} `json:"customBehaviors"`
}

func outputTableCPCodes(cpcodes *edgegrid.PropertyAPICPCodes) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)

	fmt.Fprintln(w, fmt.Sprint("# ID \t ProductIDs \t Created \t Name"))
	for _, singleCPcoode := range cpcodes.Cpcodes.Items {
		fmt.Fprintln(w, fmt.Sprintf("%s \t %s \t %s \t %s", singleCPcoode.CpcodeID, singleCPcoode.ProductIds, singleCPcoode.CreatedDate, singleCPcoode.CpcodeName))
	}

	w.Flush()

}

func outputTableProducts(products *edgegrid.PropertyAPIProducts) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)

	fmt.Fprintln(w, fmt.Sprint("# ID \t Name"))
	for _, singleProduct := range products.Products.Items {
		fmt.Fprintln(w, fmt.Sprintf("%s \t %s", singleProduct.ProductID, singleProduct.ProductName))
	}

	w.Flush()

}

func outputTableGroups(groups *edgegrid.PropertyAPIGroups) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)

	fmt.Fprintln(w, fmt.Sprint("# ID \t ContractID \t Name"))
	for _, singleGroup := range groups.Groups.Items {
		fmt.Fprintln(w, fmt.Sprintf("%s \t %s \t %s", singleGroup.GroupID, singleGroup.ContractIds, singleGroup.GroupName))
	}

	w.Flush()

}

func outputTableContracts(contracts *edgegrid.PropertyAPIContracts) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', 0)

	fmt.Fprintln(w, fmt.Sprint("# ID \t Name"))
	for _, singleContract := range contracts.Contracts.Items {
		fmt.Fprintln(w, fmt.Sprintf("%s \t %s", singleContract.ContractID, singleContract.ContractTypeName))
	}

	w.Flush()

}

// OutputJSON displays output of query for alerts in JSON format
//
// output
func OutputJSON(input interface{}) {
	b, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
