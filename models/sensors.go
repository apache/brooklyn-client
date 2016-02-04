package models

type SensorSummary struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Links       map[string]URI `json:"links"`
	Type        string         `json:"type"`
}
