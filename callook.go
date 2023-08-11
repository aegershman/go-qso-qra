package main

type CallookResponse struct {
	Status    string       `json:"status"`
	Type      string       `json:"type"`
	Current   CallookInfo  `json:"current"`
	Previous  CallookInfo  `json:"previous"`
	Trustee   TrusteeInfo  `json:"trustee"`
	Name      string       `json:"name"`
	Address   AddressInfo  `json:"address"`
	Location  LocationInfo `json:"location"`
	OtherInfo OtherInfo    `json:"otherInfo"`
}

type CallookInfo struct {
	Callsign  string `json:"callsign"`
	OperClass string `json:"operClass"`
}

type TrusteeInfo struct {
	Callsign string `json:"callsign"`
	Name     string `json:"name"`
}

type AddressInfo struct {
	Line1 string `json:"line1"`
	Line2 string `json:"line2"`
	Attn  string `json:"attn"`
}

type LocationInfo struct {
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	Gridsquare string `json:"gridsquare"`
}

type OtherInfo struct {
	GrantDate      string `json:"grantDate"`
	ExpiryDate     string `json:"expiryDate"`
	LastActionDate string `json:"lastActionDate"`
	Frn            string `json:"frn"`
	UlsUrl         string `json:"ulsUrl"`
}

func createDisplayInfoCallook(info CallookResponse) DisplayInfo {
	return DisplayInfo{
		Callsign:  info.Current.Callsign,
		Class:     info.Current.OperClass,
		Expires:   info.OtherInfo.ExpiryDate,
		Status:    info.Status,
		Grid:      info.Location.Gridsquare,
		Lat:       info.Location.Latitude,
		Lon:       info.Location.Longitude,
		FName:     "",
		MI:        "",
		Name:      info.Name,
		Suffix:    "",
		AddrLine1: info.Address.Line1,
		AddrLine2: info.Address.Line2,
		State:     "",
		Zip:       "",
		Country:   "",
	}
}
