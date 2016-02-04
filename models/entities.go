package models

type EntitySummary struct {
	CatalogItemId string         `json:"catalogItemId"`
	Name          string         `json:"name"`
	Links         map[string]URI `json:"links"`
	Id            string         `json:"id"`
	Type          string         `json:"type"`
}
