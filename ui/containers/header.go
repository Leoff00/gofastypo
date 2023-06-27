package containers

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	mutMin      string = "01"
	stopCounter        = make(chan bool, 1)
)

func beginCounter(activate bool, secTimer *widget.Label) {
	if activate {
		stopCounter <- activate
		return
	} else {
		go func() {
			for i := 59; i >= 0; i-- {
				select {
				case <-stopCounter:
					return
				default:
					secTimer.SetText(strconv.Itoa(i))
					time.Sleep(time.Second)
				}
			}
		}()
	}
}

func chooser(selectorStr string) string {
	var sv string
	if selectorStr == "1 minute" {
		sv = "01"
	} else if selectorStr == "2 minutes" {
		sv = "02"
	} else if selectorStr == "3 minutes" {
		sv = "03"
	}
	return sv
}

func HeaderContainer() *fyne.Container {
	minOptions := make([]string, 0, 3)
	minOptions = append(minOptions, "1 minute", "2 minutes", "3 minutes")

	minTimer := widget.NewLabel("00")
	separator := widget.NewLabel(":")
	secTimer := widget.NewLabel("00")
	stopMsg := widget.NewLabel("")

	optionsSel := widget.NewSelect(minOptions, func(s string) {
		mutMin = chooser(s)
		minTimer.SetText(mutMin)
	})

	optionsSel.PlaceHolder = "Select the minutage:"

	stopBtn := widget.NewButton("Stop!", func() {
		stopMsg.Show()
		stopMsg.SetText("stopped...")
		beginCounter(true, secTimer)
		optionsSel.Enable()
	})
	startBtn := widget.NewButton("Start!", func() {
		beginCounter(false, secTimer)
		stopMsg.Hide()
		optionsSel.Disable()
	})

	timer := container.NewHBox(minTimer, separator, secTimer, stopMsg)
	rowContainer := container.New(layout.NewGridLayout(4), optionsSel, startBtn, stopBtn, timer)
	return container.New(layout.NewGridLayoutWithRows(3), rowContainer)

}
