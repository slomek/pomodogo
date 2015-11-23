package notifications

import (
	"github.com/0xAX/notificator"
	"path/filepath"
	"os"
	"fmt"
)

var notify = notificator.New(notificator.Options{
	DefaultIcon: "icon/default.png",
	AppName:     "PomodoGo",
})

func showMessage(message string) {
	notify.Push("PomodoGo!", message, getImagePath("tomato.png"), notificator.UR_CRITICAL)
}

func getImagePath(fileName string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return fileName
	}
	return fmt.Sprintf("%s/%s", dir, fileName)
}

// PomodoroFinishNotification displays system notification about finished Pomodoro period.
func PomodoroFinishNotification() {
	showMessage("Your pomodoGo has finished! Time for a break!")
}
