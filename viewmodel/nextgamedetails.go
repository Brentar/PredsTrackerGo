package viewmodel

import (
	"encoding/json"
)

type NextGameDetails struct {
	GameType       string
	OpponentID     int
	Opponent       string
	OpponentWins   int
	OpponentLosses int
	OpponentOT     int
	Venue          string
	City           string
	GameDate       string
}

func (d NextGameDetails) UnmarshalJSON(bytes []byte) error {
	m := make(map[string]string)
	json.Unmarshal(bytes, m)
	return nil
}
