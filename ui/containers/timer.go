package containers

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	mutMin      = 1
	stopCounter = make(chan bool)
	minOptions  = make([]string, 0, 3)

	counter = widget.NewLabel("")
	stopMsg = widget.NewLabel("")

	Duration time.Duration
)

func beginCounter(shouldStop bool, min time.Duration) {
	Duration = min * time.Minute
	if !shouldStop {
		go func() {
			for Duration >= 0 {
				select {
				case <-stopCounter:
					TxtArea.Enable()
					return
				default:
					counter.SetText(Duration.String())
					Duration -= time.Second
					time.Sleep(time.Second)
					if Duration <= 0 {
						TxtArea.Disable()
					}
				}
			}
		}()
	} else {
		stopCounter <- shouldStop
		return
	}
}

func chooser(selectorStr string) int {
	var sv time.Duration
	if selectorStr == "1 minute" {
		sv = 1
	} else if selectorStr == "2 minutes" {
		sv = 2
	} else if selectorStr == "3 minutes" {
		sv = 3
	}
	return int(sv)
}

func HeaderContainer() *fyne.Container {
	minOptions = append(minOptions, "1 minute", "2 minutes", "3 minutes")
	optionsSel := widget.NewSelect(minOptions, func(s string) {
		mutMin = chooser(s)
	})

	optionsSel.PlaceHolder = "Select the minutage:"

	stopBtn := widget.NewButton("Stop!", func() {
		stopMsg.Show()
		stopMsg.SetText("stopped...")
		optionsSel.Enable()
		StopTyping(TxtArea)
		beginCounter(true, time.Duration(mutMin))
	})
	startBtn := widget.NewButton("Start!", func() {
		beginCounter(false, time.Duration(mutMin))
		StartTyping(TxtArea)
		stopMsg.Hide()
		optionsSel.Disable()
	})

	timer := container.NewHBox(counter, stopMsg)
	rowContainer := container.New(layout.NewGridLayout(4), optionsSel, startBtn, stopBtn, timer)
	return container.New(layout.NewGridLayoutWithRows(3), rowContainer)

}
