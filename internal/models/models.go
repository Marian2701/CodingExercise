package models

type Game struct {
	Id        uint32
	HomeTeam  Countries
	HomeScore uint
	AwayTeam  Countries
	AwayScore uint
}
