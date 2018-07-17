package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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
	fmt.Println("File server start at http://127.0.0.1:" + fmt.Sprintf("%d", port))
	err := http.ListenAndServe(":"+fmt.Sprintf("%d", port), http.FileServer(http.Dir(PATH)))

	if err != nil {
		fmt.Println("File server start error: ", err)
		fmt.Println("Reset port ...")

		port++
		return run()
	}

	return nil
}
