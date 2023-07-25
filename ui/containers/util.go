package containers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/leoff00/gofastypo/db"
)

func parseMetric(m []db.MetricModel) []string {
	var res []string
	for i, v := range m {
		res = append(res, fmt.Sprintf("wpm #%d: %.0f words/minute", (i+1), v.Metric))
	}
	return res
}

func CalculateWordsPerMinute(wordCount int, startTime time.Time) float64 {
	elapsedMin := time.Since(startTime).Minutes()
	if elapsedMin > 0 {
		wordsPerMinute := float64(wordCount) / elapsedMin
		return wordsPerMinute
	}
	return 0
}

func ExecuteStopTyping() []string {
	v, _ := strconv.Atoi(VisualMetric)
	mu.Lock()
	db.PersistData(float64(v))
	mu.Unlock()
	TxtArea.Disable()
	TxtArea.FocusGained()
	TxtArea.SetText("")
	return parseMetric(db.GetData())
}
