package nhl

import "predsTrackerGo/nhl/models"

type HomeOrAway struct {
	LeagueRecord struct {
		Wins   int    `json:"wins"`
		Losses int    `json:"losses"`
		Ot     int    `json:"ot"`
		Type   string `json:"type"`
	} `json:"leagueRecord"`
	Score int `json:"score"`
	Team  models.ScheduleTeam
}
