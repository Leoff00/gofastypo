package containers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var mutMin string = "01"

func beginCounter() *widget.Label {
	secLeft := binding.NewString()
	go func() {
		for i := 59; i >= 0; i-- {
			if err := secLeft.Set(strconv.Itoa(i)); err != nil {
				log.Fatal("cannot attribute the var in counter", err)
			}
			fmt.Println(secLeft)
			time.Sleep(time.Second)
		}
	}()
	return widget.NewLabelWithData(secLeft)
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

	optionsSel := widget.NewSelect(minOptions, func(s string) {
		mutMin = chooser(s)
		minTimer.SetText(mutMin)
	})
	optionsSel.PlaceHolder = "Select the minutage:"

	startBtn := widget.NewButton("Start!", func() {
		secLeft := beginCounter()
		secTimer.SetText(secLeft.Text)
		optionsSel.Disable()
	})

	timer := container.NewHBox(minTimer, separator, secTimer)

	rowContainer := container.NewHBox(optionsSel, startBtn, timer)
	return container.NewVBox(rowContainer)

}
