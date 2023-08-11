package main

type HamDBResponse struct {
	APIResponse struct {
		Version  string       `json:"version"`
		Callsign CallsignInfo `json:"callsign"`
		Messages struct {
			Status string `json:"status"`
		} `json:"messages"`
	} `json:"hamdb"`
}

type CallsignInfo struct {
	Call    string `json:"call"`
	Class   string `json:"class"`
	Expires string `json:"expires"`
	Status  string `json:"status"`
	Grid    string `json:"grid"`
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
	FName   string `json:"fname"`
	MI      string `json:"mi"`
	Name    string `json:"name"`
	Suffix  string `json:"suffix"`
	Addr1   string `json:"addr1"`
	Addr2   string `json:"addr2"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

func createDisplayInfoHamDB(info CallsignInfo) DisplayInfo {
	return DisplayInfo{
		Callsign:  info.Call,
		Class:     info.Class,
		Expires:   info.Expires,
		Status:    info.Status,
		Grid:      info.Grid,
		Lat:       info.Lat,
		Lon:       info.Lon,
		FName:     info.FName,
		MI:        info.MI,
		Name:      info.Name,
		Suffix:    info.Suffix,
		AddrLine1: info.Addr1,
		AddrLine2: info.Addr2,
		State:     info.State,
		Zip:       info.Zip,
		Country:   info.Country,
	}
}
