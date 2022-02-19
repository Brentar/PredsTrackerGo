package model

import "time"

type ScheduleResponse struct {
	Copyright    string         `json:"copyright"`
	TotalItems   int            `json:"totalItems"`
	TotalEvents  int            `json:"totalEvents"`
	TotalGames   int            `json:"totalGames"`
	TotalMatches int            `json:"totalMatches"`
	Dates        []ScheduleDate `json:"dates"'`
	MetaData     struct {
		TimeStamp string `json:"timeStamp"`
	} `json:"metaData"`
	Wait int `json:"wait"`
}

type ScheduleDate struct {
	Date         string         `json:"date"`
	TotalItems   int            `json:"totalItems"`
	TotalEvents  int            `json:"totalEvents"`
	TotalGames   int            `json:"totalGames"`
	Games        []ScheduleGame `json:"games"`
	TotalMatches int            `json:"totalMatches"`
	Events       []interface{}  `json:"events"`
	Matches      []interface{}  `json:"matches"`
}

type ScheduleGame struct {
	GamePk   int       `json:"gamePk"`
	Link     string    `json:"link"`
	GameType string    `json:"gameType"`
	Season   string    `json:"season"`
	GameDate time.Time `json:"gameDate"`
	Status   struct {
		AbstractGameState string `json:"abstractGameState"`
		CodedGameState    string `json:"codedGameState"`
		DetailedState     string `json:"detailedState"`
		StatusCode        string `json:"statusCode"`
		StartTimeTBD      bool   `json:"startTimeTBD"`
	} `json:"status"`
	Teams struct {
		Away struct {
			LeagueRecord struct {
				Wins   int    `json:"wins"`
				Losses int    `json:"losses"`
				Ot     int    `json:"ot"`
				Type   string `json:"type"`
			} `json:"leagueRecord"`
			Score int `json:"score"`
			Team  struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
		} `json:"away"`
		Home struct {
			LeagueRecord struct {
				Wins   int    `json:"wins"`
				Losses int    `json:"losses"`
				Ot     int    `json:"ot"`
				Type   string `json:"type"`
			} `json:"leagueRecord"`
			Score int `json:"score"`
			Team  struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
		} `json:"home"`
	} `json:"teams"`
	Venue struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"venue"`
	Content struct {
		Link string `json:"link"`
	} `json:"content"`
}
