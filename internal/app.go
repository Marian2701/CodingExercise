package internal

import (
	"errors"
	"github.com/Marian2701/CodingExercise/internal/models"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type App struct {
	store  ScoreBaseStoring
	board  GameBoard
	Server *http.Server
	logger *log.Logger
}

func NewApp(store ScoreBaseStoring, board GameBoard) *App {
	return &App{
		store:  store,
		board:  board,
		Server: nil,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

type PageData struct {
	Countries        []models.Countries
	ActiveMatches    []*models.Game
	CompletedMatches []*models.Game
}

// InitRoutes initializes HTTP routes for handling match selection, match updates, and game completion.
func (a *App) InitRoutes() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Countries:        models.AllCountries,
			ActiveMatches:    a.board.GetGames(),
			CompletedMatches: a.store.GetGames(),
		}

		tmpl, err := template.New("index").Parse(`
			<!DOCTYPE html>
			<html lang="ru">
			<head>
				<meta charset="UTF-8">
				<title>Matches</title>
			</head>
			<body>
				<h1>Selection of countries for the match</h1>
				<form method="post" action="/start_game">
					<label for="country1">First country:</label>
					<select name="country1">
						{{range .Countries}}
							<option value="{{.}}">{{.}}</option>
						{{end}}
					</select>
					<label for="country2">Second country:</label>
					<select name="country2">
						{{range .Countries}}
							<option value="{{.}}">{{.}}</option>
						{{end}}
					</select>
					<button type="submit">Start a match</button>
				</form>

				<h2>Active matches</h2>
				<ul>
					{{range .ActiveMatches}}
						<li>
							{{.HomeTeam}} - {{.AwayTeam}} | {{.HomeScore}} : {{.AwayScore}}
							<form method="post" action="/update_score">
								<input type="hidden" name="matchIndex" value="{{.Id}}">
								<input type="number" name="score1" value="{{.HomeScore}}" min="0">
								<input type="number" name="score2" value="{{.AwayScore}}" min="0">
								<button type="submit">Update the result</button>
							</form>
							<form method="post" action="/end_game">
								<input type="hidden" name="matchIndex" value="{{.Id}}">
								<button type="submit">Finish match</button>
							</form>
						</li>
					{{end}}
				</ul>

				<h2>Completed matches</h2>
				<ul>
					{{range .CompletedMatches}}
						<li>{{.HomeTeam}} - {{.AwayTeam}} | {{.HomeScore}} : {{.AwayScore}}</li>
					{{end}}
				</ul>
			</body>
			</html>
		`)

		if err != nil {
			a.logger.Println("failed to parse template: ", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			a.logger.Println("failed to execute template data: ", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/start_game", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		if err := a.board.StartGame(r.FormValue("country1"), r.FormValue("country2")); err != nil {
			if errors.Is(err, models.ErrInvalidCountry) {
				a.logger.Println("invalid country from request: ", err)
				http.Error(w, "Invalid country", http.StatusBadRequest)
				return
			} else {
				a.logger.Println("failed to init game: ", err)
				http.Error(w, "internal error", http.StatusInternalServerError)
				return
			}
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	mux.HandleFunc("/end_game", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		id, err := strconv.Atoi(r.FormValue("matchIndex"))
		if err != nil {
			a.logger.Println("failed to get id from request: ", err)
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}

		game, err := a.board.RemoveGame(uint32(id))
		if err != nil {
			if errors.Is(err, models.ErrGameNotFound) {
				a.logger.Println("invalid id from request: ", err)
				http.Error(w, "Invalid id", http.StatusBadRequest)
				return
			} else {
				a.logger.Println("failed to remove game from scoreBoard: ", err)
				http.Error(w, "internal error", http.StatusInternalServerError)
				return
			}
		}

		a.store.Insert(game)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	mux.HandleFunc("/update_score", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		id, err := strconv.Atoi(r.FormValue("matchIndex"))
		if err != nil {
			a.logger.Println("failed to get id from request: ", err)
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return
		}

		homeScore, err := strconv.Atoi(r.FormValue("score1"))
		if err != nil {
			a.logger.Println("failed to get home score from request: ", err)
			http.Error(w, "Invalid home score", http.StatusBadRequest)
			return
		}

		awayScore, err := strconv.Atoi(r.FormValue("score2"))
		if err != nil {
			a.logger.Println("failed to get away score from request: ", err)
			http.Error(w, "Invalid away score", http.StatusBadRequest)
			return
		}

		if err := a.board.UpdateGame(uint32(id), uint(homeScore), uint(awayScore)); err != nil {
			if errors.Is(err, models.ErrGameNotFound) {
				a.logger.Println("invalid id from request: ", err)
				http.Error(w, "Invalid id", http.StatusBadRequest)
				return
			} else {
				a.logger.Println("failed to update game on scoreBoard: ", err)
				http.Error(w, "internal error", http.StatusInternalServerError)
				return
			}
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	})

	a.Server = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}

func (a *App) RunServer() {
	if err := a.Server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
