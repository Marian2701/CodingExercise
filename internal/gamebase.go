package internal

import (
	"github.com/Marian2701/CodingExercise/internal/models"
	"sync"
)

type ScoreBase struct {
	Root *GameNode
	lock sync.RWMutex
}

type GameNode struct {
	Value *models.Game
	Left  *GameNode
	Right *GameNode
}

func NewScoreBase() *ScoreBase {
	return &ScoreBase{
		Root: nil,
		lock: sync.RWMutex{},
	}
}

type ScoreBaseStoring interface {
	Insert(value *models.Game)
	GetGames() []*models.Game
}

func (x *ScoreBase) Insert(value *models.Game) {
	// TODO
	panic("implement me")
}

func (x *ScoreBase) GetGames() []*models.Game {
	// TODO
	panic("implement me")
}
