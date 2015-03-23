package models

type Tree struct {
	Name         string `json:"name"`
	Id           string `json:"id"`
	Type         string `json:"type"`
	ServiceUp    bool   `json:"serviceUp"`
	ServiceState string `json:"serviceState"`
	Children     []Tree `json:"children"`
}

type TaskSummary struct {
	SubmitTimeUtc     int64                       `json:"submitTimeUtc"`
	EndTimeUtc        int64                       `json:"endTimeUtc"`
	IsCancelled       bool                        `json:"isCancelled"`
	CurrentStatus     string                      `json:"currentStatus"`
	BlockingTask      LinkWithMetadata            `json:"blockingTask"`
	DisplayName       string                      `json:"displayName"`
	Streams           map[string]LinkWithMetadata `json:"streams"`
	Description       string                      `json:"description"`
	EntityId          string                      `json:"entityId"`
	EntityDisplayName string                      `json:"entityDisplayName"`
	Error             bool                        `json:"error"`
	SubmittedByTask   LinkWithMetadata            `json:"submittedByTask"`
	Result            interface{}                 `json:"result"`
	IsError           bool                        `json:"isError"`
	DetailedStatus    string                      `json:"detailedStatus"`
	Children          []LinkWithMetadata          `json:"children"`
	BlockingDetails   string                      `json:"blockingDetails"`
	Cancelled         bool                        `json:"cancelled"`
	Links             map[string]URI              `json:"links"`
	Id                string                      `json:"id"`
	StartTimeUtc      int64                       `json:"startTimeUtc"`
}

type ApplicationSummary struct {
	Links  map[string]URI  `json:"links"`
	Id     string          `json:"id"`
	Spec   ApplicationSpec `json:"spec"`
	Status Status          `json:"status"`
}

type ApplicationSpec struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Locations []string `json:"locations"`
}

type Status string

type EntitySummary struct {
	CatalogItemId string         `json:"links"`
	Name          string         `json:"name"`
	Links         map[string]URI `json:"links"`
	Id            string         `json:"id"`
	Type          string         `json:"type"`
}

type LinkWithMetadata struct {
}

type URI string
