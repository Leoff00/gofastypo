package containers

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/leoff00/gofastypo/gen"
)

var (
	g           = gen.GeneratePhrases{}
	phrase, arr = g.Generate()
	msgs        = widget.NewLabel(phrase)
	txtArea     = widget.NewMultiLineEntry()
)

func capture(input string) {
	checkType := strings.Contains(phrase, input)
	checkLen := len(input) == len(phrase)
	zeroLen := len(arr) == 0

	if checkType {
		for i := 0; i < len(arr); i++ {
			if checkLen {
				txtArea.SetText("")
				copy(arr[i:], arr[i+1:])
				arr = arr[:len(arr)-1]
				phrase = arr[i]
				msgs.SetText(phrase)
				if zeroLen {
					msgs.SetText("")
				}
				return
			}
		}
	}
}

func TextAreaContainer() *fyne.Container {
	txtArea.SetPlaceHolder("Type here...")
	txtArea.MultiLine = true
	txtArea.Wrapping = fyne.TextWrapWord

	txtArea.OnChanged = capture

	txtAreaC := container.NewMax(txtArea)
	handler := container.New(layout.NewGridLayoutWithRows(2), msgs, txtAreaC)
	return container.NewCenter(handler)
}
