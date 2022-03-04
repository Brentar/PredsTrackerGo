package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"predsTrackerGo/nhl"
	"predsTrackerGo/repo"
	"predsTrackerGo/viewmodel"
)

func viewRoster(w http.ResponseWriter, r *http.Request) {
	roster, err := repo.GetRoster()

	tmpl, err := template.ParseFiles("view/roster.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	tmpl.Execute(w, roster)
}

func viewTimeOnIce(w http.ResponseWriter, r *http.Request) {
	timeOnIce, err := repo.GetTimeOnIce()

	tmpl, err := template.ParseFiles("view/time-on-ice.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	if err = tmpl.Execute(w, timeOnIce); err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
}

func playerInfo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/playerInfo.html")
	if err != nil {
		http.Error(w, "failed to load player information", http.StatusInternalServerError)
	}

	playerId := r.URL.Query().Get("id")
	playerInfo, _ := repo.GetPlayerInfo(playerId)

	tmpl.Execute(w, playerInfo)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	game, err := repo.GetNextScheduledGame()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	var homeOrAway nhl.HomeOrAway
	if game.Teams.Home.Team.Id == 18 {
		homeOrAway = game.Teams.Away
	} else {
		homeOrAway = game.Teams.Home
	}

	gameVm := viewmodel.NextGameDetails{
		GameType:       game.GameType,
		Opponent:       homeOrAway.Team.Name,
		OpponentWins:   homeOrAway.LeagueRecord.Wins,
		OpponentLosses: homeOrAway.LeagueRecord.Losses,
		OpponentID:     homeOrAway.Team.Id,
		Venue:          game.Venue.Name,
		GameDate:       game.Date,
	}

	tmpl, err := template.ParseFiles("view/index.html")
	tmpl.Execute(w, gameVm)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/roster", viewRoster)
	http.HandleFunc("/playerInfo", playerInfo)
	http.HandleFunc("/timeOnIce", viewTimeOnIce)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
