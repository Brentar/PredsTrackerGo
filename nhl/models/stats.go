package models

type Stats struct {
	SkaterStats struct {
		TimeOnIce            string `json:"timeOnIce"`
		Assists              int    `json:"assists"`
		Goals                int    `json:"goals"`
		Shots                int    `json:"shots"`
		Hits                 int    `json:"hits"`
		PowerPlayGoals       int    `json:"powerPlayGoals"`
		PowerPlayAssists     int    `json:"powerPlayAssists"`
		PenaltyMinutes       int    `json:"penaltyMinutes"`
		FaceOffWins          int    `json:"faceOffWins"`
		FaceoffTaken         int    `json:"faceoffTaken"`
		Takeaways            int    `json:"takeaways"`
		Giveaways            int    `json:"giveaways"`
		ShortHandedGoals     int    `json:"shortHandedGoals"`
		ShortHandedAssists   int    `json:"shortHandedAssists"`
		Blocked              int    `json:"blocked"`
		PlusMinus            int    `json:"plusMinus"`
		EvenTimeOnIce        string `json:"evenTimeOnIce"`
		PowerPlayTimeOnIce   string `json:"powerPlayTimeOnIce"`
		ShortHandedTimeOnIce string `json:"shortHandedTimeOnIce"`
	} `json:"skaterStats"`
}
