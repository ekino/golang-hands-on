package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	type Talk struct {
		Name        string                 `json:"name,omitempty"`
		Description string                 `json:"description,omitempty"`
		Private     bool                   `json:"-"`
		Place       int                    `json:"place,,string"`
		Properties  map[string]interface{} `json:"properties,omitempty"`
	}

	talk := &Talk{
		Name: "Marshalling with go",
		Private: true,
		Place: 200,
		Properties: map[string]interface{}{
			"retro": true,
			"foo": 1,
		},
	}

	data, _ := json.Marshal(talk)

	fmt.Printf("Marshal: %s\n", data)

	source := []byte(`{"name":"Marshalling with go","place":"100","Private":false,"properties":{"retro":false,"bar":"extra"}}`)

	json.Unmarshal(source, talk)

	fmt.Printf("Unmarshal: %+v", talk)
}