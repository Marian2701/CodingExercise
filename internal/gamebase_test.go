package internal

import (
	"github.com/Marian2701/CodingExercise/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScoreBase_Insert(t *testing.T) {
	tests := []struct {
		name string
		game *models.Game
		want []*models.Game
	}{
		{
			name: "Inserting first game",
			game: &models.Game{Id: 0, HomeTeam: models.Morocco, AwayTeam: models.Canada, HomeScore: 0, AwayScore: 5},
			want: []*models.Game{{Id: 0, HomeTeam: models.Morocco, AwayTeam: models.Canada, HomeScore: 0, AwayScore: 5}},
		},
		{
			name: "Inserting second game",
			game: &models.Game{Id: 1, HomeTeam: models.Spain, AwayTeam: models.Brazil, HomeScore: 10, AwayScore: 2},
			want: []*models.Game{
				{Id: 1, HomeTeam: models.Spain, AwayTeam: models.Brazil, HomeScore: 10, AwayScore: 2},
				{Id: 0, HomeTeam: models.Morocco, AwayTeam: models.Canada, HomeScore: 0, AwayScore: 5},
			},
		},
		{
			name: "Inserting third game",
			game: &models.Game{Id: 2, HomeTeam: models.Germany, AwayTeam: models.France, HomeScore: 2, AwayScore: 2},
			want: []*models.Game{
				{Id: 1, HomeTeam: models.Spain, AwayTeam: models.Brazil, HomeScore: 10, AwayScore: 2},
				{Id: 0, HomeTeam: models.Morocco, AwayTeam: models.Canada, HomeScore: 0, AwayScore: 5},
				{Id: 2, HomeTeam: models.Germany, AwayTeam: models.France, HomeScore: 2, AwayScore: 2},
			},
		},
		{
			name: "Inserting fourth game",
			game: &models.Game{Id: 3, HomeTeam: models.USA, AwayTeam: models.Italy, HomeScore: 6, AwayScore: 6},
			want: []*models.Game{
				{Id: 3, HomeTeam: models.USA, AwayTeam: models.Italy, HomeScore: 6, AwayScore: 6},
				{Id: 1, HomeTeam: models.Spain, AwayTeam: models.Brazil, HomeScore: 10, AwayScore: 2},
				{Id: 0, HomeTeam: models.Morocco, AwayTeam: models.Canada, HomeScore: 0, AwayScore: 5},
				{Id: 2, HomeTeam: models.Germany, AwayTeam: models.France, HomeScore: 2, AwayScore: 2},
			},
		},
		{
			name: "Inserting fifth game",
			game: &models.Game{Id: 4, HomeTeam: models.Argentina, AwayTeam: models.Australia, HomeScore: 3, AwayScore: 1},
			want: []*models.Game{
				{Id: 3, HomeTeam: models.USA, AwayTeam: models.Italy, HomeScore: 6, AwayScore: 6},
				{Id: 1, HomeTeam: models.Spain, AwayTeam: models.Brazil, HomeScore: 10, AwayScore: 2},
				{Id: 0, HomeTeam: models.Morocco, AwayTeam: models.Canada, HomeScore: 0, AwayScore: 5},
				{Id: 4, HomeTeam: models.Argentina, AwayTeam: models.Australia, HomeScore: 3, AwayScore: 1},
				{Id: 2, HomeTeam: models.Germany, AwayTeam: models.France, HomeScore: 2, AwayScore: 2},
			},
		},
	}

	sb := NewScoreBase()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sb.Insert(tt.game)
			got := sb.GetGames()
			assert.Equal(t, len(tt.want), len(got))
			for i, g := range got {
				assert.Equal(t, tt.want[i].Id, g.Id)
				assert.Equal(t, true, GameEquality(tt.want[i], g))
			}
		})
	}
}

func GameEquality(game1, game2 *models.Game) bool {
	return game1.Id == game2.Id &&
		game1.HomeTeam == game2.HomeTeam &&
		game1.AwayTeam == game2.AwayTeam &&
		game1.HomeScore == game2.HomeScore &&
		game1.AwayScore == game2.AwayScore
}
