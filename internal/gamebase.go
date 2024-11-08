package internal

import (
	"github.com/Marian2701/CodingExercise/internal/models"
	"sync"
)

// ScoreBase represents the base structure for storing game scores with a root node of the BTS(binary search tree) and a sync RWMutex.
// RWMutex is used for concurrency-safe work with this structure.
// BTS was chosen as the storage structure because the main required operations are as follows
// 1. Inserting an element, the complexity of this operation in BTS is O(log(N))
// 2. Getting all elements in sorted form, the complexity of this operation in BTS is O(N),
// since the structure and the insertion operation imply storing the data in sorted form.
type ScoreBase struct {
	Root *GameNode
	lock sync.RWMutex
}

// GameNode represents a node in BTS(binary search tree) with a pointer to a game entity and left and right child nodes.
type GameNode struct {
	Value *models.Game
	Left  *GameNode
	Right *GameNode
}

// NewScoreBase returns a new instance of ScoreBase with a nil root node and a sync RWMutex.
func NewScoreBase() *ScoreBase {
	return &ScoreBase{
		Root: nil,
		lock: sync.RWMutex{},
	}
}

// ScoreBaseStoring defines methods for storing game scores, including inserting a new game and getting all stored games.
type ScoreBaseStoring interface {
	Insert(value *models.Game)
	GetGames() []*models.Game
}

// Insert adds a new game node with the provided game data to the binary search tree.
// If the root node is nil, the new node becomes the root; otherwise, it is inserted following the BST rules.
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

// insertNode adds a newNode to the binary search tree starting from the given node following the BST rules.
// If the newNode score is greater than or equal to the node's score, it is inserted to the left; otherwise, to the right.
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

// GetGames returns a slice of all games stored in the binary search tree.
// It applies an in-order traversal starting from the root node to collect and return all games.
// If the root is nil, an empty slice is returned.
func (x *ScoreBase) GetGames() []*models.Game {
	x.lock.RLock()
	defer x.lock.RUnlock()

	var result []*models.Game
	if x.Root != nil {
		result = inOrderTraverse(x.Root, result)
	}
	return result
}

// inOrderTraverse performs in-order traversal on a binary search tree starting
// from the given node and returns a slice of games in sorted order.
// If the node is nil, an empty slice is returned.
func inOrderTraverse(node *GameNode, result []*models.Game) []*models.Game {
	if node != nil {
		result = inOrderTraverse(node.Left, result)
		result = append(result, node.Value)
		result = inOrderTraverse(node.Right, result)
	}
	return result
}
