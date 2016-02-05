package commands

import (
	"encoding/json"
)

func stringRepresentation(value interface{}) (string, error) {
	var result string;
	switch value.(type) {
	case string:
		result = value.(string)  // use string value as-is
	default:
		json, err := json.Marshal(value)
		if err != nil {
			return "", err
		}
		result = string(json)   // return JSON text representation of value object
	}
	return result, nil;
}

