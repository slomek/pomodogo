package main

import (
	"flag"
	"fmt"
	"strings"
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
		min, sec := formatSeconds(left)
		progress := progressBar(total, left)
		fmt.Printf("\r%02d:%02d %s", min, sec, progress)
		left -= 1
	}
	notifications.ShowMessage("End")
}

func formatSeconds(seconds int) (int, int) {
	m := seconds / 60
	s := seconds % 60
	return m, s
}

func progressBar(total int, left int) string {
	all := screenutils.ColumnsCount() - 6

	leftPercent := all * left / total
	complPercent := all - leftPercent

	completedBlocks := strings.Repeat("█", complPercent)
	leftBlocks := strings.Repeat("▒", leftPercent)

	return completedBlocks + leftBlocks
}




