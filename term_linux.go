package dump

import (
	"syscall"
)

const ioctlReadTermios = syscall.TCGETS

type Termios syscall.Termios
