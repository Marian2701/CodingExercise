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
	x.lock.Lock()
	defer x.lock.Unlock()

	newNode := &GameNode{Value: value}
	if x.Root == nil {
		x.Root = newNode
	} else {
		insertNode(x.Root, newNode)
	}
}

func insertNode(node, newNode *GameNode) {
	if newNode.Value.GetScore() > node.Value.GetScore() || newNode.Value.GetScore() == node.Value.GetScore() {
		if node.Left == nil {
			node.Left = newNode
		} else {
			insertNode(node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNode(node.Right, newNode)
		}
	}
}

func (x *ScoreBase) GetGames() []*models.Game {
	x.lock.RLock()
	defer x.lock.RUnlock()

	var result []*models.Game
	if x.Root != nil {
		result = inOrderTraverse(x.Root, result)
	}
	return result
}

func inOrderTraverse(node *GameNode, result []*models.Game) []*models.Game {
	if node != nil {
		result = inOrderTraverse(node.Left, result)
		result = append(result, node.Value)
		result = inOrderTraverse(node.Right, result)
	}
	return result
}
