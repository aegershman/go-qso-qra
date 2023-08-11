package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Provide a callsign. Usage:")
		fmt.Println("qso kd0pmb")
		os.Exit(1)
	}
	callsign := os.Args[1]
	callsign = strings.ToUpper(callsign)

	apiFlag := flag.String("api", "hamdb", "API endpoint (hamdb or callook)")
	flag.Parse()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("ERROR: Could not determine user's home directory: %s\n", err)
		os.Exit(1)
	}

	cacheDir := filepath.Join(homeDir, ".ham_cache")
	err = os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		fmt.Printf("ERROR: Could not create cache directory: %s\n", err)
		os.Exit(1)
	}

	cacheFilePath := filepath.Join(cacheDir, callsign+".json")

	var cachedData CachedData
	if cachedDataBytes, err := os.ReadFile(cacheFilePath); err == nil {
		err = json.Unmarshal(cachedDataBytes, &cachedData)
		if err != nil {
			fmt.Printf("ERROR: Could not unmarshal cached data: %s\n", err)
			os.Exit(1)
		} else if cachedData.API == *apiFlag {
			printCachedResponse(cachedData.Response, *apiFlag)
			return
		}
	}

	var apiEndpoint string
	switch *apiFlag {
	case "hamdb":
		apiEndpoint = fmt.Sprintf("https://api.hamdb.org/v1/%s/json/hamdb", callsign)
	case "callook":
		apiEndpoint = fmt.Sprintf("https://callook.info/%s/json", callsign)
	default:
		fmt.Println("Invalid API endpoint specified")
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

	printCachedResponse(responseBody, *apiFlag)
}
