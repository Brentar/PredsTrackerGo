package model

type RosterEntry struct {
	Person struct {
		Id       int    `json:"id"`
		FullName string `json:"fullName"`
		Link     string `json:"link"`
	} `json:"person"`

	JerseyNumber string `json:"jerseyNumber"`

	Position struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Abbreviation string `json:"abbreviation"`
	} `json:"position"`
}
