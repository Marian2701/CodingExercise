package internal

import (
	"github.com/Marian2701/CodingExercise/internal/models"
	"sync"
	"sync/atomic"
)

// ScoreBoard represents a scoreboard containing games using a concurrent-safe map.
// The choice of this data structure is due to the fact that the basic operations in this case are
// 1. Inserting an element, the complexity of this operation in map is O(1)
// 2. Deleting an element, the complexity of this action in map is O(1)
// 3. Updating an element, the complexity of this action in the map is O(1)
// 4. Retrieving all elements, the complexity of this action in map is O(N)
// From this it was concluded that it was well suited for storing active matches
type ScoreBoard struct {
	Games  *sync.Map
	nextId uint32
}

// NewScoreBoard initializes a new scoreboard with games stored in a concurrent-safe map.
func NewScoreBoard() *ScoreBoard {
	return &ScoreBoard{
		Games: &sync.Map{},
	}
}

// GameBoard defines methods for managing games on a game board.
// It allows starting a game, removing a game, updating scores, and getting all games.
type GameBoard interface {
	StartGame(homeTeam, awayTeam string) error
	RemoveGame(id uint32) (*models.Game, error)
	UpdateGame(id uint32, homeScore, awayScore uint) error
	GetGames() []*models.Game
}

const (
	// beginHomeScore is the initial score value for the home team in a game on the scoreboard.
	beginHomeScore = 0
	// beginAwayScore is the initial score value for the away team in a game on the scoreboard.
	beginAwayScore = 0
)

// StartGame initializes a new game with the provided home and away teams, assigns initial scores, and increments the game ID.
func (x *ScoreBoard) StartGame(homeTeam, awayTeam string) error {
	id := atomic.AddUint32(&x.nextId, 1)
	homeTeamCountry := models.GetCountryFromString(homeTeam)
	if homeTeamCountry == models.NotACountry {
		return models.ErrInvalidCountry
	}
	awayTeamCountry := models.GetCountryFromString(awayTeam)
	if awayTeamCountry == models.NotACountry {
		return models.ErrInvalidCountry
	}

	x.Games.Store(id, &models.Game{
		Id:        id,
		HomeTeam:  homeTeamCountry,
		AwayTeam:  awayTeamCountry,
		HomeScore: beginHomeScore,
		AwayScore: beginAwayScore,
	})

	return nil
}

// RemoveGame removes a game from the scoreboard by the provided ID, returning the removed game.
func (x *ScoreBoard) RemoveGame(id uint32) (*models.Game, error) {
	game, ok := x.Games.LoadAndDelete(id)
	if !ok {
		return nil, models.ErrGameNotFound
	}
	return game.(*models.Game), nil
}

// UpdateGame finds the game with the provided ID in the scoreboard, sets the home and away scores
// for the game, and swaps the updated game back into the scoreboard.
func (x *ScoreBoard) UpdateGame(id uint32, homeScore, awayScore uint) error {
	gameMap, ok := x.Games.Load(id)
	if !ok {
		return models.ErrGameNotFound
	}

	game := gameMap.(*models.Game)

	game.SetHomeScore(homeScore).SetAwayScore(awayScore)

	x.Games.Swap(id, game)

	return nil
}

// GetGames retrieves all games stored in the scoreboard and returns them as a slice of Game pointers.
func (x *ScoreBoard) GetGames() []*models.Game {
	var result []*models.Game
	x.Games.Range(func(key, value interface{}) bool {
		result = append(result, value.(*models.Game))
		return true
	})
	return result
}
