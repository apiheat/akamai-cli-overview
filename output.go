package main

import (
	"time"
)

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
