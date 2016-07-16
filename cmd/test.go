package main

import (
	"fmt"
	"github.com/2qif49lt/dump"
	"io"
	"net/http"
	"os"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	i := 0
	fmt.Println(10 / i)

	fmt.Fprintln(os.Stderr, dump.IsTerminal(), os.Stderr.Fd(), dump.IsFile())
	io.WriteString(w, fmt.Sprintf(`hello world %v %v  %v %v`, dump.IsTerminal(), os.Stderr.Fd(), dump.IsFile(), dump.IsNeedDump))

}
func main() {
	http.HandleFunc("/hello", HelloServer)
	http.ListenAndServe(":8888", nil)
}

// test cmdline
/*
./cmd
./cmd &
./cmd 2>&1
./cmd 2>&1 &
nohup ./cmd >1.txt 2>&1 &

nohup ./cmd >/dev/null 2>&1 &    dump files
*/
