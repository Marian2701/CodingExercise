package internal

import (
	"log"
	"net/http"
	"os"
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
