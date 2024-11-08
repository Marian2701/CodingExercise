package internal

import (
	"github.com/Marian2701/CodingExercise/internal/models"
	"sync"
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

func (x *ScoreBoard) StartGame(homeTeam, awayTeam string) error {
	// TODO
	panic("implement me")
}

func (x *ScoreBoard) RemoveGame(id uint32) (*models.Game, error) {
	// TODO
	panic("implement me")
}

func (x *ScoreBoard) UpdateGame(id uint32, homeScore, awayScore uint) error {
	// TODO
	panic("implement me")
}

func (x *ScoreBoard) GetGames() []*models.Game {
	// TODO
	panic("implement me")
}
