package dump

import (
	"os"
)

var IsNeedDump = !IsTerminal() && !IsFile()

func init() {
	if IsNeedDump {
		f, err := os.Create("panic.dump")
		if err == nil {
			redirectStderr(f)
		}
	}

}
