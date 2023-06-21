package containers

import (
	"strconv"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	mutMin      string = "01"
	stopCounter        = make(chan bool)
	mu          sync.Mutex
)

// ! and know one way to stop the go routine exec.
func beginCounter(activate bool, secTimer *widget.Label) {
	go func() {
		for i := 59; i >= 0; i-- {
			select {
			case <-stopCounter:
				return
			default:
				time.Sleep(time.Second)
				secTimer.SetText(strconv.Itoa(i))
			}
		}
	}()
	if activate {
		stopCounter <- activate
		return
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
		mu.Lock()
		stopCounter <- true
		stopMsg.Show()
		stopMsg.SetText("stopped...")
		beginCounter(true, secTimer)
		optionsSel.Enable()
		mu.Unlock()
	})
	startBtn := widget.NewButton("Start!", func() {
		mu.Lock()
		beginCounter(false, secTimer)
		stopMsg.Hide()
		optionsSel.Disable()
		mu.Unlock()
	})

	timer := container.NewHBox(minTimer, separator, secTimer, stopMsg)

	rowContainer := container.NewHBox(optionsSel, startBtn, stopBtn, timer)
	return container.NewVBox(rowContainer)

}
