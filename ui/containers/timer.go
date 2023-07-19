package containers

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	stopCounter = make(chan bool)
	minOptions  = make([]string, 0, 3)
	mutMin      = 1

	optionsSel = widget.NewSelect([]string{}, func(s string) {})
	stopBtn    = widget.NewButton("Stop", func() {})
	startBtn   = widget.NewButton("Start", func() {})
	counter    = widget.NewLabel("")
	stopMsg    = widget.NewLabel("")

	Duration time.Duration
)

func setCounterMinutes(min time.Duration) time.Duration {
	return min * time.Minute
}

func changeOnStop() {
	stopMsg.Show()
	stopMsg.SetText("stopped...")
	optionsSel.Enable()
	stopBtn.Disable()
	startBtn.Enable()
	go StopTyping(TxtArea)
	beginCounter(true, time.Duration(mutMin))
}

func changeOnStart() {
	msgs.SetText(phrase)
	TxtArea.FocusGained()
	TxtArea.Enable()
	msgs.Show()
	StartTyping(TxtArea)
	beginCounter(false, time.Duration(mutMin))
	stopMsg.Hide()
	optionsSel.Disable()
	stopBtn.Enable()
	startBtn.Disable()
}

func DurationReachesZero() {
	go StopTyping(TxtArea)
	TxtArea.Disable()
	stopBtn.Disable()
	startBtn.Enable()
	optionsSel.Enable()
}

func beginCounter(shouldStop bool, min time.Duration) {
	Duration = setCounterMinutes(min)
	if !shouldStop {
		go func() {
			for Duration >= 0 {
				select {
				case <-stopCounter:
					Duration = setCounterMinutes(min)
					go StopTyping(TxtArea)
					return
				default:
					counter.SetText(Duration.String())
					Duration -= time.Second
					time.Sleep(time.Second)
					if Duration <= 0 {
						go DurationReachesZero()
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
	optionsSel = widget.NewSelect(minOptions, func(s string) {
		startBtn.Enable()
		mutMin = chooser(s)
	})

	startBtn = widget.NewButton("Start!", func() {
		changeOnStart()
	})

	stopBtn = widget.NewButton("Stop!", func() {
		changeOnStop()
	})

	optionsSel.PlaceHolder = "Select the minutage:"
	startBtn.Disable()
	stopBtn.Disable()

	timer := container.NewHBox(counter, stopMsg)
	rowContainer := container.New(layout.NewGridLayout(4), optionsSel, startBtn, stopBtn, timer)
	return container.New(layout.NewGridLayoutWithRows(3), rowContainer)

}
