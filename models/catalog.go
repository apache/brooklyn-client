package models

type CatalogItemSummary struct {
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

type CatalogPolicySummary struct {
	symbolicName string         `json:"symbolicName"`  
    version      string         `json:"version"`  
    displayName  string         `json:"name"`  
    javaType     string         `json:"javaType"`  
    planYaml     string         `json:"planYaml"`  
    description  string         `json:"description"`  
    iconUrl      string         `json:"iconUrl"`  
    deprecated   bool           `json:"deprecated"`  
    links        map[string]URI `json:"links"` 
}


type CatalogLocationSummary struct {
	
}

type CatalogEntitySummary struct {
	symbolicName string                 `json:"symbolicName"`  
    version      string                 `json:"version"`  
    displayName  string                 `json:"name"`  
    javaType     string                 `json:"javaType"`  
    planYaml     string                 `json:"planYaml"`  
    description  string                 `json:"description"`
    Config       []ConfigSummary        `json:"config"`
    Effectors    []EffectorSummary      `json:"effectors"`
    Sensors      []SensorSummary        `json:"sensors"`
	Deprecated   string                 `json:"deprecated"`
	Links        map[string]interface{} `json:"links"`
	Id           string                 `json:"id"`
	Type         string                 `json:"type"`
}
