// +build !windows

package dump

import (
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

// IsTerminal returns true if stderr's file descriptor is a terminal.
func IsTerminal() bool {
	fd := syscall.Stderr
	var termios Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(fd), ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}

func redirectStderr(f *os.File) {
	err := syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		fmt.Fprintln(os.Stderr, "dup2", err)
	} else {
		fmt.Fprintf(os.Stderr, "%s: begin dump std err.", time.Now())
	}
}
