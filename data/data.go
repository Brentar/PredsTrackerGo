package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"predsTrackerGo/model"
	"time"
)

func GetRoster() ([]model.RosterEntry, error) {
	response, err := http.Get("https://statsapi.web.nhl.com/api/v1/teams/18/roster")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	rosterResponse := model.RosterResponse{}
	if err = json.Unmarshal(responseData, &rosterResponse); err != nil {
		return nil, err
	}
	if rosterResponse.Roster == nil || len(rosterResponse.Roster) < 1 {
		return nil, errors.New("failed to load roster")
	}

	return rosterResponse.Roster, nil
}

func GetPlayerInfo(playerId string) (*model.PeopleEntry, error) {
	nhlResponse, err := http.Get("https://statsapi.web.nhl.com/api/v1/people/" + playerId)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	responseData, err := ioutil.ReadAll(nhlResponse.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var response model.PeopleResponse
	json.Unmarshal(responseData, &response)
	if response.People == nil || len(response.People) != 1 {
		return nil, errors.New("failed to load roster")
	}

	return &response.People[0], nil
}

func GetNextGameDetails() (*model.ScheduleGame, error) {
	now := time.Now()
	nextYear := now.Add(time.Duration(now.Year() + 1))
	nhlResponse, err := http.Get(fmt.Sprintf("https://statsapi.web.nhl.com/api/v1/schedule?teamId=18&startDate=%v&endDate=%v", time.Now(), nextYear))

	responseData, err := ioutil.ReadAll(nhlResponse.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var response model.ScheduleResponse
	json.Unmarshal(responseData, &response)
	if err != nil {
		return nil, err
	}
	if response.Dates == nil || len(response.Dates) < 1 {
		return nil, errors.New("failed to retrieve next game details")
	}

	return &response.Dates[0].Games[0], nil
}
