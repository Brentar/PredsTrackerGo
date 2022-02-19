package model

type RosterResponse struct {
	Copyright string        `json:"copyright"`
	Roster    []RosterEntry `json:"roster"`
}
