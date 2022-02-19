package model

type Player struct {
	Person       Person   `json:"person"`
	JerseyNumber string   `json:"jerseyNumber"`
	Position     Position `json:"position"`
}

type Position struct {
	Name string `json:"name"`
}

type Person struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
}

type Players []Player

type Response struct {
	Copyright string  `json:"copyright"`
	Players   Players `json:"people"`
}
