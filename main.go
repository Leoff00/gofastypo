package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/leoff00/gofastypo/db"
	"github.com/leoff00/gofastypo/ui"
)

func main() {
	go db.StartDB()

	fyne := app.New()
	run := ui.Init(fyne)
	run.ShowAndRun()

}
