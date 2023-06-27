package containers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/leoff00/gofastypo/gen"
)

func TextAreaContainer() *fyne.Container {
	g := gen.GeneratePhrases{}
	msgs := widget.NewLabel(g.Generate())

	txtArea := widget.NewMultiLineEntry()
	txtArea.SetPlaceHolder("Type here...")
	txtArea.MultiLine = true
	txtArea.Wrapping = fyne.TextWrapWord

	txtAreaC := container.NewMax(txtArea)
	handler := container.New(layout.NewGridLayoutWithRows(2), msgs, txtAreaC)
	return container.NewCenter(handler)
}
