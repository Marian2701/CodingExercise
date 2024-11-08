package models

type Game struct {
	Id        uint32
	HomeTeam  Countries
	HomeScore uint
	AwayTeam  Countries
	AwayScore uint
}

func (x *Game) SetHomeScore(homeScore uint) *Game {
	x.HomeScore = homeScore
	return x
}

func (x *Game) SetAwayScore(awayScore uint) *Game {
	x.AwayScore = awayScore
	return x
}

func (x *Game) GetScore() uint {
	return x.HomeScore + x.AwayScore
}
