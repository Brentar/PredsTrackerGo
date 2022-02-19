package model

type PeopleResponse struct {
	Copyright string `json:"copyright"`
	People    []PeopleEntry
}

type PeopleEntry struct {
	Id               int    `json:"id"`
	FullName         string `json:"fullName"`
	Link             string `json:"link"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	PrimaryNumber    string `json:"primaryNumber"`
	BirthDate        string `json:"birthDate"`
	CurrentAge       int    `json:"currentAge"`
	BirthCity        string `json:"birthCity"`
	BirthCountry     string `json:"birthCountry"`
	Nationality      string `json:"nationality"`
	Height           string `json:"height"`
	Weight           int    `json:"weight"`
	Active           bool   `json:"active"`
	AlternateCaptain bool   `json:"alternateCaptain"`
	Captain          bool   `json:"captain"`
	Rookie           bool   `json:"rookie"`
	ShootsCatches    string `json:"shootsCatches"`
	RosterStatus     string `json:"rosterStatus"`
	CurrentTeam      struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"currentTeam"`
	PrimaryPosition struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Abbreviation string `json:"abbreviation"`
	} `json:"primaryPosition"`
}
