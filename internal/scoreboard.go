package internal

import (
	"github.com/Marian2701/CodingExercise/internal/models"
	"sync"
	"sync/atomic"
)

type ScoreBoard struct {
	Games  *sync.Map
	nextId uint32
}

func NewScoreBoard() *ScoreBoard {
	return &ScoreBoard{
		Games: &sync.Map{},
	}
}

type GameBoard interface {
	StartGame(homeTeam, awayTeam string) error
	RemoveGame(id uint32) (*models.Game, error)
	UpdateGame(id uint32, homeScore, awayScore uint) error
	GetGames() []*models.Game
}

const (
	beginHomeScore = 0
	beginAwayScore = 0
)

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

func (x *ScoreBoard) RemoveGame(id uint32) (*models.Game, error) {
	game, ok := x.Games.LoadAndDelete(id)
	if !ok {
		return nil, models.ErrGameNotFound
	}
	return game.(*models.Game), nil
}

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

func (x *ScoreBoard) GetGames() []*models.Game {
	var result []*models.Game
	x.Games.Range(func(key, value interface{}) bool {
		result = append(result, value.(*models.Game))
		return true
	})
	return result
}
