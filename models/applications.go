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
	SubmitTimeUtc     int64                              `json:"submitTimeUtc"`
	EndTimeUtc        int64                              `json:"endTimeUtc"`
	IsCancelled       bool                               `json:"isCancelled"`
	CurrentStatus     string                             `json:"currentStatus"`
	BlockingTask      LinkTaskWithMetadata               `json:"blockingTask"`
	DisplayName       string                             `json:"displayName"`
	Streams           map[string]LinkStreamsWithMetadata `json:"streams"`
	Description       string                             `json:"description"`
	EntityId          string                             `json:"entityId"`
	EntityDisplayName string                             `json:"entityDisplayName"`
	Error             bool                               `json:"error"`
	SubmittedByTask   LinkTaskWithMetadata               `json:"submittedByTask"`
	Result            interface{}                        `json:"result"`
	IsError           bool                               `json:"isError"`
	DetailedStatus    string                             `json:"detailedStatus"`
	Children          []LinkTaskWithMetadata             `json:"children"`
	BlockingDetails   string                             `json:"blockingDetails"`
	Cancelled         bool                               `json:"cancelled"`
	Links             map[string]URI                     `json:"links"`
	Id                string                             `json:"id"`
	StartTimeUtc      int64                              `json:"startTimeUtc"`
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

type LinkWithMetadata struct {
}

type LinkStreamsWithMetadata struct {
	Link     string             `json:"link"`
	Metadata LinkStreamMetadata `json:"metadata"`
}

type LinkStreamMetadata struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	SizeText string `json:"sizeText"`
}

type LinkTaskWithMetadata struct {
	Link     string           `json:"link"`
	Metadata LinkTaskMetadata `json:"metadata"`
}

type LinkTaskMetadata struct {
	Id                string `json:"id"`
	TaskName          string `json:"taskName"`
	EntityId          string `json:"entityId"`
	EntityDisplayName string `json:"entityDisplayName"`
}

type URI string
