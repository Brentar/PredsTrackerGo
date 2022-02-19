package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"predsTrackerGo/data"
)

func viewRoster(w http.ResponseWriter, r *http.Request) {
	roster, err := data.GetRoster()

	tmpl, err := template.ParseFiles("view/roster.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	tmpl.Execute(w, roster)
}

func playerInfo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("view/playerInfo.html")
	if err != nil {
		http.Error(w, "failed to load player information", http.StatusInternalServerError)
	}

	playerId := r.URL.Query().Get("id")
	playerInfo, _ := data.GetPlayerInfo(playerId)

	tmpl.Execute(w, playerInfo)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	nextGameDetails, err := data.GetNextGameDetails()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	if nextGameDetails.Teams.Home.Team.Id == 18 {
		
	}

	tmpl, err := template.ParseFiles("index.html")

	tmpl.Execute(w, nextGameDetails)
	//http.ServeFile(w, r, "index.html")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/roster", viewRoster)
	http.HandleFunc("/playerInfo", playerInfo)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
