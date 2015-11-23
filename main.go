package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/slomek/pomodogo/notifications"
	"github.com/slomek/pomodogo/screenutils"
)

var duration int

func init() {
	flag.IntVar(&duration, "duration", 25, "Duration of your Pomodoro")
}

func main() {
	flag.Parse()

	total := duration * 60
	left := total
	for range time.Tick(time.Second) {
		if left < 0 {
			break
		}
		min, sec := screenutils.FormatSeconds(left)
		progress := screenutils.ShowProgressBar(total, left)
		fmt.Printf("\r%02d:%02d %s", min, sec, progress)
		left -= 1
	}

	notifications.PomodoroFinishNotification()
}
