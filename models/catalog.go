package models

type Application struct {
	Name         string                 `json:"name"`
	JavaType     string                 `json:"javaType"`
	SymbolicName string                 `json:"symbolicName"`
	Version      string                 `json:"version"`
	PlanYaml     string                 `json:"planYaml"`
	Description  string                 `json:"description"`
	Deprecated   string                 `json:"deprecated"`
	Links        map[string]interface{} `json:"links"`
	Id           string                 `json:"id"`
	Type         string                 `json:"type"`
}
