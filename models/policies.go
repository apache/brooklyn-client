package models

type PolicySummary struct {
	CatalogItemId string         `json:"catalogItemId"`
	Name          string         `json:"name"`
	Links         map[string]URI `json:"links"`
	Id            string         `json:"id"`
	State         Status         `json:"state"`
}
