package containers

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	mu                     sync.Mutex
	sharedGoroutineStarted bool
	storedMetric           float64
	VisualMetric           string
	startTime              time.Time
	wordCount              = 0
	data                   = []string{"No metrics yet"}
	metricView             = widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("No Metrics Yet...")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})
)

func StartTyping(entry *widget.Entry) {
	startTime = time.Now()
	mu.Lock()
	if !sharedGoroutineStarted {
		sharedGoroutineStarted = true
		mu.Unlock()

		go func() {
			for {
				mu.Lock()
				words := strings.Fields(entry.Text)
				wordCount = len(words)
				time.Sleep(time.Second)

				storedMetric = CalculateWordsPerMinute(wordCount, startTime)
				VisualMetric = strconv.Itoa(int(storedMetric))
				mu.Unlock()
				if Duration <= 0 {
					break
				}
			}
		}()
	} else {
		mu.Unlock()
	}
}

func StopTyping(txtArea *widget.Entry) {
	data = ExecuteStopTyping()
}

func MeterContainer() *fyne.Container {
	metricHeader := widget.NewLabel("Your metrics:")
	savedMetrics := container.NewVScroll(metricView)
	metricBox := container.NewVBox(metricHeader, savedMetrics)
	rows := container.New(layout.NewGridLayoutWithRows(1), metricBox)
	return container.NewCenter(rows)
}
