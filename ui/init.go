package ui

import (
	"fyne.io/fyne/v2"
	"github.com/leoff00/gofastypo/ui/containers"
)

func Init(app fyne.App) fyne.Window {
	window := app.NewWindow("Go Fastypo")
	window.Resize(fyne.NewSize(600, 400))
	window.CenterOnScreen()

	window.SetContent(containers.HeaderContainer())

	return window
}
