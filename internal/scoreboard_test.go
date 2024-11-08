package internal

import (
	"fmt"
	"github.com/Marian2701/CodingExercise/internal/models"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func getNumOfGames(scoreboard *ScoreBoard) int {
	counter := 0
	scoreboard.Games.Range(func(key, value interface{}) bool {
		counter++
		return true
	})
	return counter
}

func TestScoreBoard_StartGame(t *testing.T) {
	tests := []struct {
		name     string
		homeTeam string
		awayTeam string
	}{
		{
			name:     "Correct countries",
			homeTeam: "Argentina",
			awayTeam: "Brazil",
		},
		{
			name:     "Correct countries",
			homeTeam: "Egypt",
			awayTeam: "Germany",
		},
		{
			name:     "Correct countries",
			homeTeam: "Poland",
			awayTeam: "USA",
		},
	}

	scoreboard := NewScoreBoard()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numOfGamesBefore := getNumOfGames(scoreboard)
			err := scoreboard.StartGame(tt.homeTeam, tt.awayTeam)
			assert.NoError(t, err)
			numOfGamesAfter := getNumOfGames(scoreboard)
			if numOfGamesAfter != numOfGamesBefore+1 {
				t.Errorf("number of games after starting new game should be 1 more than before, got: %v, want: %v", numOfGamesAfter, numOfGamesBefore+1)
			}
		})
	}
}

func TestScoreBoard_StartGame_WrongCountries(t *testing.T) {
	tests := []struct {
		name     string
		homeTeam string
		awayTeam string
	}{
		{
			name:     "Invalid countries",
			homeTeam: "Norway",
			awayTeam: "Brazil",
		},
		{
			name:     "Empty countries",
			homeTeam: "",
			awayTeam: "Brazil",
		},
	}

	scoreboard := NewScoreBoard()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := scoreboard.StartGame(tt.homeTeam, tt.awayTeam)
			assert.ErrorIs(t, err, models.ErrInvalidCountry)
		})
	}
}

func TestScoreBoard_StartGame_concurrently(t *testing.T) {
	scoreboard := NewScoreBoard()
	numOfGoroutines := 1000
	var wg sync.WaitGroup
	wg.Add(numOfGoroutines)
	for i := 0; i < numOfGoroutines; i++ {
		go func() {
			defer wg.Done()
			err := scoreboard.StartGame("Australia", "Poland")
			assert.NoError(t, err)
		}()
	}
	wg.Wait()
	if numOfGames := getNumOfGames(scoreboard); numOfGames != numOfGoroutines {
		t.Errorf("number of games after starting new game concurrently should be %v, got: %v", numOfGoroutines, numOfGames)
	}
}

func TestScoreBoard_RemoveGame(t *testing.T) {
	tests := []struct {
		name string
		id   uint32
	}{
		{
			name: "RemovingGame",
			id:   1,
		},
		{
			name: "RemovingGame",
			id:   2,
		},
		{
			name: "RemovingGame",
			id:   3,
		},
		{
			name: "RemovingGame",
			id:   4,
		},
		{
			name: "RemovingGame",
			id:   5,
		},
	}

	testData := []struct {
		HomeTeam string
		AwayTeam string
	}{
		{HomeTeam: "USA", AwayTeam: "Italy"},
		{HomeTeam: "Spain", AwayTeam: "Brazil"},
		{HomeTeam: "Morocco", AwayTeam: "Canada"},
		{HomeTeam: "Argentina", AwayTeam: "Australia"},
		{HomeTeam: "Germany", AwayTeam: "France"},
	}

	scoreboard := NewScoreBoard()

	for _, datum := range testData {
		err := scoreboard.StartGame(datum.HomeTeam, datum.AwayTeam)
		assert.NoError(t, err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := scoreboard.RemoveGame(tt.id)
			assert.NoError(t, err)
			assert.NotNil(t, game)
			assert.Equal(t, tt.id, game.Id)
		})
	}

	assert.Equal(t, 0, len(scoreboard.GetGames()))
}

func TestScoreBoard_RemoveGame_WrongId(t *testing.T) {
	tests := []struct {
		name string
		id   uint32
	}{
		{
			name: "Not Presented",
			id:   99,
		},
	}

	testData := []struct {
		HomeTeam string
		AwayTeam string
	}{
		{HomeTeam: "USA", AwayTeam: "Italy"},
		{HomeTeam: "Spain", AwayTeam: "Brazil"},
		{HomeTeam: "Morocco", AwayTeam: "Canada"},
		{HomeTeam: "Argentina", AwayTeam: "Australia"},
		{HomeTeam: "Germany", AwayTeam: "France"},
	}

	scoreboard := NewScoreBoard()

	for _, datum := range testData {
		err := scoreboard.StartGame(datum.HomeTeam, datum.AwayTeam)
		assert.NoError(t, err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			game, err := scoreboard.RemoveGame(tt.id)
			assert.ErrorIs(t, err, models.ErrGameNotFound)
			assert.Nil(t, game)
		})
	}

	fmt.Println(len(scoreboard.GetGames()), len(testData))

	assert.Equal(t, len(testData), len(scoreboard.GetGames()))
}

func TestScoreBoard_GetGames(t *testing.T) {
	testData := []struct {
		HomeTeam string
		AwayTeam string
	}{
		{HomeTeam: "USA", AwayTeam: "Italy"},
		{HomeTeam: "Spain", AwayTeam: "Brazil"},
		{HomeTeam: "Morocco", AwayTeam: "Canada"},
		{HomeTeam: "Argentina", AwayTeam: "Australia"},
		{HomeTeam: "Germany", AwayTeam: "France"},
	}

	scoreboard := NewScoreBoard()

	for i, datum := range testData {
		err := scoreboard.StartGame(datum.HomeTeam, datum.AwayTeam)
		assert.NoError(t, err)
		assert.Equal(t, i+1, getNumOfGames(scoreboard))
	}

	assert.Equal(t, getNumOfGames(scoreboard), len(scoreboard.GetGames()))
}

func TestScoreBoard_UpdateGame(t *testing.T) {
	tests := []struct {
		name         string
		id           uint32
		newValueHome uint
		newValueAway uint
		resultHome   uint
		resultAway   uint
	}{
		{
			name:         "Add value to zero game score",
			id:           1,
			newValueHome: 0,
			newValueAway: 1,
			resultHome:   0,
			resultAway:   1,
		},
		{
			name:         "Add value to zero game score",
			id:           2,
			newValueHome: 2,
			newValueAway: 3,
			resultHome:   2,
			resultAway:   3,
		},
		{
			name:         "Add value to not zero game score",
			id:           1,
			newValueHome: 2,
			newValueAway: 3,
			resultHome:   2,
			resultAway:   3,
		},
		{
			name:         "Add value to not zero game score",
			id:           2,
			newValueHome: 1,
			newValueAway: 3,
			resultHome:   1,
			resultAway:   3,
		},
	}

	testData := []struct {
		HomeTeam string
		AwayTeam string
	}{
		{HomeTeam: "USA", AwayTeam: "Italy"},
		{HomeTeam: "Spain", AwayTeam: "Brazil"},
	}

	scoreboard := NewScoreBoard()

	for _, datum := range testData {
		err := scoreboard.StartGame(datum.HomeTeam, datum.AwayTeam)
		assert.NoError(t, err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := scoreboard.UpdateGame(tt.id, tt.newValueHome, tt.newValueAway)
			assert.NoError(t, err)
			games := scoreboard.GetGames()
			g := &models.Game{}
			for _, game := range games {
				if game.Id == tt.id {
					g = game
				}
			}
			assert.Equal(t, tt.resultHome, g.HomeScore)
			assert.Equal(t, tt.resultAway, g.AwayScore)
		})
	}
}

func TestScoreBoard_UpdateGame_WrongId(t *testing.T) {
	tests := []struct {
		name string
		id   uint32
	}{
		{
			name: "Add value to zero game score",
			id:   99,
		},
	}

	testData := []struct {
		HomeTeam string
		AwayTeam string
	}{
		{HomeTeam: "USA", AwayTeam: "Italy"},
		{HomeTeam: "Spain", AwayTeam: "Brazil"},
	}

	scoreboard := NewScoreBoard()

	for _, datum := range testData {
		err := scoreboard.StartGame(datum.HomeTeam, datum.AwayTeam)
		assert.NoError(t, err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := scoreboard.UpdateGame(tt.id, 0, 0)
			assert.ErrorIs(t, err, models.ErrGameNotFound)
		})
	}
}
