package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/leoff00/gofastypo/ui"
)

func main() {
	fyne := app.New()
	run := ui.Init(fyne)
	run.ShowAndRun()

}

// g := gen.GenerateArrays{}

// fmt.Println(g.Generate())
