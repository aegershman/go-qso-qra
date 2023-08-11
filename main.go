package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type CachedData struct {
	API      string          `json:"api"`
	Response json.RawMessage `json:"response"`
}

func main() {
	callsignFlag := flag.String("callsign", "", "ham radio call sign string")
	apiFlag := flag.String("api", "hamdb", "API endpoint (hamdb or callook)")
	verboseFlag := flag.Bool("verbose", false, "Print additional response info")
	flag.Parse()

	if *callsignFlag == "" {
		fmt.Println("Please provide a callsign using -callsign flag.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("ERROR: Could not determine user's home directory: %s\n", err)
		os.Exit(1)
	}

	cacheDir := filepath.Join(homeDir, ".hamdb_cache")
	err = os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		fmt.Printf("ERROR: Could not create cache directory: %s\n", err)
		os.Exit(1)
	}

	cacheFilePath := filepath.Join(cacheDir, *callsignFlag+".json")

	var cachedData CachedData
	if cachedDataBytes, err := os.ReadFile(cacheFilePath); err == nil {
		err = json.Unmarshal(cachedDataBytes, &cachedData)
		if err != nil {
			fmt.Printf("ERROR: Could not unmarshal cached data: %s\n", err)
			os.Exit(1)
		} else if cachedData.API == *apiFlag {
			if *verboseFlag {
				fmt.Println("Found cached data. Using cached response:")
			}
			printCachedResponse(cachedData.Response, *apiFlag)
			return
		}
	}

	var apiEndpoint string
	switch *apiFlag {
	case "hamdb":
		apiEndpoint = fmt.Sprintf("https://api.hamdb.org/v1/%s/json/hamdb", *callsignFlag)
	case "callook":
		apiEndpoint = fmt.Sprintf("https://callook.info/%s/json", *callsignFlag)
	default:
		fmt.Println("Invalid API specified. Supported values: hamdb or callook")
		os.Exit(1)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Get(apiEndpoint)
	if err != nil {
		fmt.Printf("ERROR: Could not send HTTP request to API endpoint: %s\n", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("ERROR: Non-200 HTTP response code retrieved from API endpoint")
		os.Exit(1)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("ERROR: Could not read response body: %s\n", err)
		os.Exit(1)
	}

	cachedData.API = *apiFlag
	cachedData.Response = responseBody
	cachedDataBytes, err := json.Marshal(cachedData)
	if err != nil {
		fmt.Printf("ERROR: Could not marshal cached data: %s\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(cacheFilePath, cachedDataBytes, os.ModePerm)
	if err != nil {
		fmt.Printf("ERROR: Could not write to cache file: %s\n", err)
	}

	if *verboseFlag {
		fmt.Println("API Response:")
	}

	printCachedResponse(responseBody, *apiFlag)
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
