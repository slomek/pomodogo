package screenutils

import (
	"syscall"
	"unsafe"
	"strings"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func columnsCount() int {
	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		panic(errno)
	}
	return int(ws.Col)
}

// ProgressBar is responsible for preparing pretty progress bar as ASCII-art.
func ProgressBar(total int, left int) string {
	all := columnsCount() - 6

	leftPercent := all * left / total
	complPercent := all - leftPercent

	completedBlocks := strings.Repeat("█", complPercent)
	leftBlocks := strings.Repeat("▒", leftPercent)

	return completedBlocks + leftBlocks
}

// FormatSeconds splits given number of seconds into a pair of minutes and seconds.
func FormatSeconds(seconds int) (int, int) {
	m := seconds / 60
	s := seconds % 60
	return m, s
}