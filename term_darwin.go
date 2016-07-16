package dump

import (
	"syscall"
)

const ioctlReadTermios = syscall.TIOCGETA

type Termios syscall.Termios
