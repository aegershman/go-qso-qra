package main

import (
	"fmt"
)

type DisplayInfo struct {
	Callsign  string
	Class     string
	Expires   string
	Status    string
	Grid      string
	Lat       string
	Lon       string
	FName     string
	MI        string
	Name      string
	Suffix    string
	AddrLine1 string
	AddrLine2 string
	State     string
	Zip       string
	Country   string
}

func printDisplayInfo(info DisplayInfo) {
	fmt.Printf("Callsign: %s\n", info.Callsign)
	fmt.Printf("Class: %s\n", info.Class)
	fmt.Printf("Expires: %s\n", info.Expires)
	fmt.Printf("Status: %s\n", info.Status)
	fmt.Printf("Grid: %s\n", info.Grid)
	fmt.Printf("Latitude: %s\n", info.Lat)
	fmt.Printf("Longitude: %s\n", info.Lon)
	fmt.Printf("First Name: %s\n", info.FName)
	fmt.Printf("Middle Initial: %s\n", info.MI)
	fmt.Printf("Last Name: %s\n", info.Name)
	fmt.Printf("Suffix: %s\n", info.Suffix)
	fmt.Printf("Address 1: %s\n", info.AddrLine1)
	fmt.Printf("Address 2: %s\n", info.AddrLine2)
	fmt.Printf("State: %s\n", info.State)
	fmt.Printf("Zip: %s\n", info.Zip)
	fmt.Printf("Country: %s\n", info.Country)

	// Create a Google Maps link with the address fields
	if info.Lat != "" && info.Lon != "" {
		googleMapsLink := fmt.Sprintf("https://www.google.com/maps?q=%s,%s", info.Lat, info.Lon)
		fmt.Printf("Google Maps Link: %s\n", googleMapsLink)
	}
}
