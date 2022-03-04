package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"predsTrackerGo/model"
	"predsTrackerGo/nhl"
	"predsTrackerGo/nhl/models"
	"strconv"
	"time"
)

type predsIdType int

var predsId predsIdType = 18

func (p predsIdType) txt() string {
	return string(p)
	return strconv.Itoa(int(p)) //strconv.Itoa(p)
}

func GetRoster() ([]model.RosterEntry, error) {
	rosterResponse := model.RosterResponse{}
	url := fmt.Sprintf("/teams/%v/roster", predsId)
	if err := nhl.ParseResponse(url, &rosterResponse); err != nil {
		return nil, err
	}
	if rosterResponse.Roster == nil || len(rosterResponse.Roster) < 1 {
		return nil, errors.New("failed to load roster")
	}

	return rosterResponse.Roster, nil
}

func GetPlayerInfo(playerId string) (*model.PeopleEntry, error) {
	response := model.PeopleResponse{}
	if err := nhl.ParseResponse("people/"+playerId, &response); err != nil {
		return nil, err
	}
	if response.People == nil || len(response.People) != 1 {
		return nil, errors.New("failed to load roster")
	}

	return &response.People[0], nil
}

func getPreviousGameId() (int, error) {
	url := getPreviousGameUrl()
	response := models.ScheduleResponse{}
	if err := nhl.ParseResponse(url, &response); err != nil {
		return 0, err
	}

	count := len(response.Dates)
	for i := count - 1; i >= 0; i-- {
		if response.Dates[i].Games[0].Status.DetailedState == "Final" {
			return response.Dates[i].Games[0].GamePk, nil
		} else {
			continue
		}
	}

	return 0, errors.New("failed to get previous game id")
}

func getPreviousGameUrl() string {
	now := time.Now()
	nowText := now.Format("2006-01-02")
	previousMonthText := now.AddDate(0, -1, 0).Format("2006-01-02")
	return fmt.Sprintf("schedule?teamId=%v&startDate=%v&endDate=%v", predsId, previousMonthText, nowText)
}

func getNextGameUrl() string {
	now := time.Now()
	nowText := now.Format("2006-01-02")
	nextMonthText := now.AddDate(0, 1, 0).Format("2006-01-02")
	return fmt.Sprintf("schedule?teamId=%v&startDate=%v&endDate=%v", predsId.txt(), nowText, nextMonthText)
}

func GetNextScheduledGame() (*models.ScheduledGame, error) {
	response := models.ScheduleResponse{}
	if err := nhl.ParseResponse(getNextGameUrl(), &response); err != nil {
		return nil, err
	}

	if response.Dates == nil || len(response.Dates) < 1 {
		return nil, errors.New("failed to retrieve next game details")
	}

	for _, date := range response.Dates {
		if date.Games[0].Status.DetailedState == "Final" {
			continue
		} else {
			gameDate, err := getGameDate(date.Date)
			if err != nil {
				return nil, err
			}
			date.Games[0].Date = gameDate
			return &date.Games[0], nil
		}
	}

	return nil, errors.New("no games found within the next month")
}

func getGameDate(date string) (string, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v, %v %v, %v", t.Weekday().String(), t.Month().String(), t.Day(), t.Year()), nil
}

func GetTimeOnIce() ([]model.TimeOnIce, error) {
	id, err := getPreviousGameId()
	if err != nil {
		return nil, err
	}

	boxScore := models.BoxScore{}
	url := fmt.Sprintf("game/%v/boxscore", strconv.Itoa(id))
	if err := nhl.ParseResponse(url, &boxScore); err != nil {
		return nil, err
	}
	//if boxScore.Teams.Home.Team.Id == strconv.Itoa(18) {

	//}
	var timesOnIce []model.TimeOnIce
	for _, element := range boxScore.Teams.Home.Players {
		bytes, err := json.Marshal(element)
		bsp := models.BoxScorePlayer{}
		if err = json.Unmarshal(bytes, &bsp); err != nil {
			timeOnIce := model.TimeOnIce{
				PlayerName: bsp.Person.FullName,
				TimeOnIce:  bsp.Stats.SkaterStats.TimeOnIce,
			}
			timesOnIce = append(timesOnIce, timeOnIce)
		}
	}

	return timesOnIce, nil
}
