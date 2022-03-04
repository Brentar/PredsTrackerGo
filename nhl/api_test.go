package nhl

import (
	"encoding/json"
	"predsTrackerGo/nhl/models"
	"testing"
)

func TestParseBoxScore(t *testing.T) {
	boxScore := models.BoxScore{}
	if err := ParseResponse("game/2021020511/boxscore", &boxScore); err != nil {
		t.Error("failed to parse response")
	}

	for _, element := range boxScore.Teams.Home.Players {
		bytes, err := json.Marshal(element)
		bsp := models.BoxScorePlayer{}
		if err = json.Unmarshal(bytes, &bsp); err != nil {
			t.Error("failed to unmarshal")
		}
	}
}
