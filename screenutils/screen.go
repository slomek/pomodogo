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

func ShowProgressBar(total int, left int) string {
	all := columnsCount() - 6

	leftPercent := all * left / total
	complPercent := all - leftPercent

	completedBlocks := strings.Repeat("█", complPercent)
	leftBlocks := strings.Repeat("▒", leftPercent)

	return completedBlocks + leftBlocks
}

func FormatSeconds(seconds int) (int, int) {
	m := seconds / 60
	s := seconds % 60
	return m, s
}