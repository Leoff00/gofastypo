package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"github.com/leoff00/gofastypo/ui/containers"
)

func Init(app fyne.App) fyne.Window {

	custom := &CustomTheme{}
	app.Settings().SetTheme(custom)

	window := app.NewWindow("Go Fastypo")
	window.Resize(fyne.NewSize(600, 400))
	window.CenterOnScreen()

	c := containers.Composer(containers.HeaderContainer(), containers.TextAreaContainer())

	composer := container.New(layout.NewGridLayoutWithRows(3), c.Header, c.Textarea)
	window.SetContent(composer)
	window.SetIcon(theme.ComputerIcon())
	return window
}
