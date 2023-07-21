package containers

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/leoff00/gofastypo/db"
)

var (
	mu                     sync.Mutex
	startTime              time.Time
	ticker                 = time.NewTicker(time.Second)
	sharedGoroutineStarted bool
	storedMetric           float64
	visualMetric           string
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

func parseMetric(m []db.MetricModel) []string {
	var res []string
	for i, v := range m {
		res = append(res, fmt.Sprintf("wpm #%d: %.0f words/minute", (i+1), v.Metric))
	}
	return res
}

func calculateWordsPerMinute(wordCount int, startTime time.Time) float64 {
	elapsedMin := time.Since(startTime).Minutes()
	if elapsedMin > 0 {
		wordsPerMinute := float64(wordCount) / elapsedMin
		return wordsPerMinute
	}
	return 0
}

//! NEED TO THINK IN ONE SOLUTION TO FIX THE STORED METRIC
//! AT START START TYPING FUNCTION.

func StartTyping(entry *widget.Entry) {
	mu.Lock()
	if !sharedGoroutineStarted {
		sharedGoroutineStarted = true
		mu.Unlock()
		startTime = time.Now()

		go func() {
			for range ticker.C {
				mu.Lock()
				words := strings.Fields(entry.Text)
				wordCount = len(words)
				time.Sleep(time.Second)

				storedMetric = calculateWordsPerMinute(wordCount, startTime)
				visualMetric = strconv.Itoa(int(storedMetric))
				fmt.Println(storedMetric)
				fmt.Println(startTime)
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
	ticker.Stop()
	v, _ := strconv.Atoi(visualMetric)
	go db.PersistData(float64(v))
	data = parseMetric(db.GetData())
	defer txtArea.SetText("")
	txtArea.Disable()
	txtArea.FocusGained()
}

func MeterContainer() *fyne.Container {
	metricHeader := widget.NewLabel("Your metrics:")
	savedMetrics := container.NewVScroll(metricView)
	metricBox := container.NewVBox(metricHeader, savedMetrics)
	rows := container.New(layout.NewGridLayoutWithRows(1), metricBox)
	return container.NewCenter(rows)
}
