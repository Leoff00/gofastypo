package containers

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func HeaderContainer() *fyne.Container {
	minOptions := make([]string, 0, 3)
	minOptions = append(minOptions, "1 minute", "2 minutes", "3 minutes")

	optionsSel := widget.NewSelect(minOptions, func(s string) { fmt.Println(s) })
	optionsSel.PlaceHolder = "Select the minutage: "

	startBtn := widget.NewButton("Start!", func() {})

	rowContainer := container.NewHBox(optionsSel, startBtn)

	return container.NewVBox(rowContainer)
}
