package utils

import (
	"fmt"
	"strconv"
	"time"
)

type Minutage struct {
	Minutes, Seconds string
}

func MinutageTicker() chan string {
	tick := make(chan string)
	go func() {
		ticker := time.NewTicker(time.Second)
		for t := range ticker.C {
			formatted := fmt.Sprintf("%v", strconv.Itoa(t.Local().Second()))
			tick <- formatted
		}
	}()

	return tick
}
