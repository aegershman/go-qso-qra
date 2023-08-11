package main

import (
	"encoding/json"
	"fmt"
)

type CachedData struct {
	API      string          `json:"api"`
	Response json.RawMessage `json:"response"`
}

func printCachedResponse(data json.RawMessage, api string) {
	var info DisplayInfo
	switch api {
	case "hamdb":
		var result HamDBResponse
		err := json.Unmarshal(data, &result)
		if err != nil {
			fmt.Printf("ERROR: Could not unmarshal JSON response: %s\n", err)
			return
		}
		info = createDisplayInfoHamDB(result.APIResponse.Callsign)
	case "callook":
		var result CallookResponse
		err := json.Unmarshal(data, &result)
		if err != nil {
			fmt.Printf("ERROR: Could not unmarshal JSON response: %s\n", err)
			return
		}
		info = createDisplayInfoCallook(result)
	default:
		fmt.Printf("ERROR: Unsupported API: %s\n", api)
		return
	}

	printDisplayInfo(info)
}
