package notifications

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/0xAX/notificator"
)

var imagePath string

func init() {
	flag.StringVar(&imagePath, "icon", "~/pomogodo/icon.png", "location of notification icon")
}

var notify = notificator.New(notificator.Options{
	DefaultIcon: getImagePath("tomato.png"),
	AppName:     "PomodoGo",
})

func showMessage(message string) {
	notify.Push("PomodoGo!", message, imagePath, notificator.UR_CRITICAL)
}

func getImagePath(fileName string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return fileName
	}
	return fmt.Sprintf("%s/%s", dir, fileName)
}

// Finish displays system notification about finished Pomodoro period.
func Finish() {
	showMessage("Your pomodoGo has finished! Time for a break!")
}
