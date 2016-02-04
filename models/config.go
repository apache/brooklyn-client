package models

type ConfigSummary struct {
	Reconfigurable bool                `json:"reconfigurable"`
	PossibleValues []map[string]string `json:"possibleValues"`
	DefaultValue   interface{}         `json:"defaultValue"`
	Name           string              `json:"name"`
	Description    string              `json:"description"`
	Links          map[string]URI      `json:"links"`
	Label          string              `json:"label"`
	Priority       float64             `json:"priority"`
	Type           string              `json:"type"`
}
