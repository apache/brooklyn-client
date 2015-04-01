package models

type AccessSummary struct {
	Links map[string]URI `json:"links"`
	LocationProvisioningAllowed bool  `json:"locationProvisioningAllowed"`
}