package containers

import (
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/leoff00/gofastypo/gen"
)

var (
	g           = gen.GeneratePhrases{}
	phrase, arr = g.Generate()
	msgs        = widget.NewLabel(phrase)
	TxtArea     = widget.NewMultiLineEntry()
)

func capture(input string) {
	checkType := strings.Contains(phrase, input)
	checkLen := len(input) == len(phrase)
	regen := len(arr) == 2

	if checkType {
		for i := 0; i < len(arr); i++ {
			if checkLen {
				TxtArea.SetText("")
				copy(arr[i:], arr[i+1:])
				arr = arr[:len(arr)-1]
				time.Sleep(1 * time.Second)
				phrase = arr[i]
				msgs.SetText(phrase)
				if regen {
					phrase, arr = g.Generate()
					phrase = arr[i]
					msgs.SetText(phrase)
					return
				}
				return
			}
		}
		return

	}
}

func TextAreaContainer() *fyne.Container {
	TxtArea.SetPlaceHolder("Type here...")
	TxtArea.MultiLine = true
	TxtArea.Wrapping = fyne.TextWrapWord
	TxtArea.Disable()

	TxtArea.OnChanged = capture

	handler := container.NewVBox(container.NewCenter(msgs), TxtArea)
	return handler
}
