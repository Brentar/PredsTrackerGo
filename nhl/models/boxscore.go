package models

type BoxScore struct {
	Copyright string `json:"copyright"`
	Teams     struct {
		Away struct {
			Team struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
			TeamStats struct {
				TeamSkaterStats struct {
					Goals                  int     `json:"goals"`
					Pim                    int     `json:"pim"`
					Shots                  int     `json:"shots"`
					PowerPlayPercentage    string  `json:"powerPlayPercentage"`
					PowerPlayGoals         float64 `json:"powerPlayGoals"`
					PowerPlayOpportunities float64 `json:"powerPlayOpportunities"`
					FaceOffWinPercentage   string  `json:"faceOffWinPercentage"`
					Blocked                int     `json:"blocked"`
					Takeaways              int     `json:"takeaways"`
					Giveaways              int     `json:"giveaways"`
					Hits                   int     `json:"hits"`
				} `json:"teamSkaterStats"`
			} `json:"teamStats"`
			Players   map[string]interface{}
			Goalies   []int `json:"goalies"`
			Skaters   []int `json:"skaters"`
			OnIce     []int `json:"onIce"`
			OnIcePlus []struct {
				PlayerId      int `json:"playerId"`
				ShiftDuration int `json:"shiftDuration"`
				Stamina       int `json:"stamina"`
			} `json:"onIcePlus"`
			Scratches  []int         `json:"scratches"`
			PenaltyBox []interface{} `json:"penaltyBox"`
			Coaches    []struct {
				Person struct {
					FullName string `json:"fullName"`
					Link     string `json:"link"`
				} `json:"person"`
				Position struct {
					Code         string `json:"code"`
					Name         string `json:"name"`
					Type         string `json:"type"`
					Abbreviation string `json:"abbreviation"`
				} `json:"position"`
			} `json:"coaches"`
		} `json:"away"`
		Home struct {
			Team struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
				Link string `json:"link"`
			} `json:"team"`
			TeamStats struct {
				TeamSkaterStats struct {
					Goals                  int     `json:"goals"`
					Pim                    int     `json:"pim"`
					Shots                  int     `json:"shots"`
					PowerPlayPercentage    string  `json:"powerPlayPercentage"`
					PowerPlayGoals         float64 `json:"powerPlayGoals"`
					PowerPlayOpportunities float64 `json:"powerPlayOpportunities"`
					FaceOffWinPercentage   string  `json:"faceOffWinPercentage"`
					Blocked                int     `json:"blocked"`
					Takeaways              int     `json:"takeaways"`
					Giveaways              int     `json:"giveaways"`
					Hits                   int     `json:"hits"`
				} `json:"teamSkaterStats"`
			} `json:"teamStats"`
			Players   map[string]interface{}
			Goalies   []int `json:"goalies"`
			Skaters   []int `json:"skaters"`
			OnIce     []int `json:"onIce"`
			OnIcePlus []struct {
				PlayerId      int `json:"playerId"`
				ShiftDuration int `json:"shiftDuration"`
				Stamina       int `json:"stamina"`
			} `json:"onIcePlus"`
			Scratches  []int         `json:"scratches"`
			PenaltyBox []interface{} `json:"penaltyBox"`
			Coaches    []struct {
				Person struct {
					FullName string `json:"fullName"`
					Link     string `json:"link"`
				} `json:"person"`
				Position struct {
					Code         string `json:"code"`
					Name         string `json:"name"`
					Type         string `json:"type"`
					Abbreviation string `json:"abbreviation"`
				} `json:"position"`
			} `json:"coaches"`
		} `json:"home"`
	} `json:"teams"`
	Officials []struct {
		Official struct {
			Id       int    `json:"id"`
			FullName string `json:"fullName"`
			Link     string `json:"link"`
		} `json:"official"`
		OfficialType string `json:"officialType"`
	} `json:"officials"`
}
