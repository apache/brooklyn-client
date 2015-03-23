package models

type LocationSummary struct{
 	Id string `json:"id"`
	Name string `json:"name"`
    Spec string `json:"spec"`
    Type string `json:"type"`
    Config map[string]interface{} `json:"config"`
    Links map[string]URI `json:"links"`
}