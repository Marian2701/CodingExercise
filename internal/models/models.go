package models

// Game represents a game entity with an id, home and away teams, and respective scores.
type Game struct {
	Id        uint32
	HomeTeam  Countries
	HomeScore uint
	AwayTeam  Countries
	AwayScore uint
}

// SetHomeScore sets the home score of the game to the provided value and returns the updated game.
func (x *Game) SetHomeScore(homeScore uint) *Game {
	x.HomeScore = homeScore
	return x
}

// SetAwayScore sets the away score of the game to the provided value and returns the updated game.
func (x *Game) SetAwayScore(awayScore uint) *Game {
	x.AwayScore = awayScore
	return x
}

// GetScore calculates and returns the total score of the game, sum of home and away scores.
func (x *Game) GetScore() uint {
	return x.HomeScore + x.AwayScore
}
