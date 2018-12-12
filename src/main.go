package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var PATH string
var port = 9989

func main() {

	if len(os.Args) < 2 {
		PATH = "./"
	} else {
		PATH, _ = filepath.Abs(os.Args[1])
	}

	run()
}

func run() error {

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println("File server start error: ", err)
		return nil
	}

	inet := addrs[1]

	fmt.Println("File server start at " + strings.Split(inet.String(), "/")[0] + ":" + fmt.Sprintf("%d", port))
	err = http.ListenAndServe(":"+fmt.Sprintf("%d", port), http.FileServer(http.Dir(PATH)))

	if err != nil {
		fmt.Println("File server start error: ", err)
		fmt.Println("Reset port ...")

		port++
		return run()
	}

	return nil
}
