package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/slomek/pomodogo/notifications"
	"github.com/slomek/pomodogo/screenutils"
)

var duration time.Duration

func init() {
	flag.DurationVar(&duration, "duration", 25*time.Minute, "Duration of your Pomodoro")
}

func main() {
	flag.Parse()

	total := int(duration.Seconds())
	left := total
	for range time.Tick(time.Second) {
		if left < 0 {
			break
		}
		min, sec := screenutils.FormatSeconds(left)
		progress := screenutils.ProgressBar(total, left)
		fmt.Printf("\r%02d:%02d %s", min, sec, progress)
		left--
	}

	notifications.Finish()
}
