package containers

import (
	"fmt"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	wordCount           = 0
	wpm                 = widget.NewLabel("")
	wordsPerMinuteLabel string
)

func StartTyping(entry *widget.Entry) {
	go func() {
		startTime := time.Now()
		for Duration >= 0 {
			time.Sleep(time.Second)
			entryText := entry.Text
			words := strings.Fields(entryText)
			wordCount = len(words)
			wordsPerMinuteLabel = fmt.Sprintf("%.0f words/minute", calculateWordsPerMinute(&wordCount, startTime))
			wpm.SetText(wordsPerMinuteLabel)
		}
	}()

}

func StopTyping(txtArea *widget.Entry) {
	wordsPerMinuteLabel = "0"
	defaultMsg := fmt.Sprintf("%s words/minute", wordsPerMinuteLabel)
	wpm.SetText(defaultMsg)
	txtArea.SetText("")
	txtArea.Disable()
	txtArea.FocusGained()
}

func calculateWordsPerMinute(wordCount *int, startTime time.Time) float64 {
	elapsedTime := time.Since(startTime).Minutes()
	wordsPerMinute := float64(*wordCount) / elapsedTime
	return wordsPerMinute
}

func MeterContainer() *fyne.Container {
	return container.New(layout.NewCenterLayout(), wpm)
}
