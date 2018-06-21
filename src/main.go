package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var PATH string

func main() {

	if len(os.Args) < 2 {
		PATH = "./"
	} else {
		PATH, _ = filepath.Abs(os.Args[1])
	}

	fmt.Println("File server start at http://127.0.0.1:9990")
	err := http.ListenAndServe(":9990", http.FileServer(http.Dir(PATH)))

	if err != nil {
		fmt.Println("File server start error: ", err)
	}

}
