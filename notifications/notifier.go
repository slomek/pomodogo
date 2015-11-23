package notifications

import "github.com/0xAX/notificator"

var notify = notificator.New(notificator.Options{
	DefaultIcon: "icon/default.png",
	AppName:     "PomodoGo",
})

func showMessage(message string) {
	notify.Push("PomodoGo!", message, "", notificator.UR_CRITICAL)
}

func PomodoroFinishNotification() {
	showMessage("Your pomodoGo has finished! Time for a break!")
}
