package dump

import (
	"bufio"
	"os"
)

// IsFile check to see stderr whether is be redirected to a normal file.
func IsFile() bool {
	st, err := os.Stderr.Stat()
	if err != nil {
		return false
	}
	return st.Mode().IsRegular()
}

func createRedirectPipe(f *os.File) (*os.File, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}

	go func() {
		scanner := bufio.NewScanner(r)
		// func ScanPanic(data []byte, atEOF bool) (advance int, token []byte, err error)
		for scanner.Scan() {
			line := scanner.Text()
			f.WriteString(line)
		}
	}()

	return w, nil
}
