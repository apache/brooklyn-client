package models

type PolicySummary struct {
	CatalogItemId string         `json:"catalogItemId"`
	Name          string         `json:"name"`
	Links         map[string]URI `json:"links"`
	Id            string         `json:"id"`
	State         Status         `json:"state"`
}

type PolicyConfigList struct {
	Name           string         `json:"name"`
	Type           string         `json:"type"`
	DefaultValue   interface{}    `json:"defaultValue`
	Description    string         `json:"description"`
	Reconfigurable bool           `json:"reconfigurable"`
	Label          string         `json:"label"`
	Priority       int64          `json:"priority"`
	PossibleValues []interface{}  `json:"possibleValues"`
	Links          map[string]URI `json:"links"`
}
