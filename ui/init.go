package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/leoff00/gofastypo/ui/containers"
)

func Init(app fyne.App) fyne.Window {
	window := app.NewWindow("Go Fastypo")
	window.Resize(fyne.NewSize(600, 400))
	window.CenterOnScreen()

	composer := container.New(layout.NewGridLayoutWithRows(3), containers.HeaderContainer(), containers.TextAreaContainer())

	window.SetContent(composer)

	return window
}
