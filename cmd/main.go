package main

import (
	"github.com/Marian2701/CodingExercise/internal"
)

func main() {
	scoreBoard := internal.NewScoreBoard()
	scoreBase := internal.NewScoreBase()

	app := internal.NewApp(scoreBase, scoreBoard)
	app.InitRoutes()
	app.RunServer()
}
