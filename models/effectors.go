package models

type EffectorSummary struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Links       map[string]URI     `json:"links"`
	Parameters  []ParameterSummary `json:"parameters"`
	ReturnType  string             `json:"returnType"`
}

type ParameterSummary struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Description  string      `json:"description"`
	DefaultValue interface{} `json:"defaultValue"`
}
