package model

type PlayerInfo struct {
	FullName     string `json:"fullName"`
	Number       string `json:"primaryNumber"`
	BirthDate    string
	Age          string
	BirthCity    string
	BirthCountry string
	Height       string
	Weight       string
}

type PlayerInfoResponse struct {
	Copyright string       `json:"copyright"`
	People    []PlayerInfo `json:"people"`
}
