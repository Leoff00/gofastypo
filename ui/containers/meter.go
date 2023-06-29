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

		for Duration > 0 {
			time.Sleep(1 * time.Second)

			entryText := entry.Text
			words := strings.Fields(entryText)
			wordCount = len(words)
			wordsPerMinuteLabel = fmt.Sprintf("%.0f words/minute", calculateWordsPerMinute(wordCount, startTime))
			wpm.SetText(wordsPerMinuteLabel)
			fmt.Println(words)

		}
	}()
}

func StopTyping(entry *widget.Entry) {
	entry.SetText("")
	wpm.SetText(wordsPerMinuteLabel)
}

func calculateWordsPerMinute(wordCount int, startTime time.Time) float64 {
	elapsedTime := time.Since(startTime).Minutes()
	wordsPerMinute := float64(wordCount) / elapsedTime
	return wordsPerMinute
}

func MeterContainer() *fyne.Container {
	return container.New(layout.NewCenterLayout(), wpm)
}
