package notifications

import "github.com/0xAX/notificator"

var notify *notificator.Notificator

func ShowMessage(message string) {
	notify = notificator.New(notificator.Options{
		DefaultIcon: "icon/default.png",
		AppName:     "My test App",
	})

	notify.Push("title", message, "/home/user/icon.png", notificator.UR_CRITICAL)
}