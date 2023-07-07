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
	msgs        = widget.NewLabel(gen.BeginPhrase)
	TxtArea     = widget.NewMultiLineEntry()
)

func popPhrases(index int) {
	TxtArea.SetText("")
	copy(arr[index:], arr[index+1:])
	arr = arr[:len(arr)-1]
	time.Sleep(500 * time.Millisecond)
	phrase = arr[index]
	msgs.SetText(phrase)
}

func regenPhrases(index int) {
	phrase, arr = g.Generate()
	phrase = arr[index]
	msgs.SetText(phrase)
}

func textAreaLogic(input string) {
	checkType := strings.Contains(phrase, input)
	checkLen := len(input) == len(phrase)
	shouldRegen := len(arr) == 2

	if checkType {
		for i := 0; i < len(arr); i++ {
			if checkLen {
				popPhrases(i)
				if shouldRegen {
					regenPhrases(i)
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

	TxtArea.OnChanged = textAreaLogic

	handler := container.NewVBox(container.NewCenter(msgs), TxtArea)
	return handler
}
