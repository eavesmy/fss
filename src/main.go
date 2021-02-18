package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"strings"
)

var Path string
var Port int

func main() {

	flag.IntVar(&Port, "p", 9989, "-p 80")
	flag.StringVar(&Path, "d", "./", "-d ./")
	flag.Parse()

	Path, _ = filepath.Abs(Path)

	run()
}

func run() error {

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println("File server start error: ", err)
		return nil
	}

	inet := addrs[1]

	fmt.Println("File server start at "+Path, strings.Split(inet.String(), "/")[0]+":"+fmt.Sprintf("%d", Port))
	err = http.ListenAndServe(":"+fmt.Sprintf("%d", Port), http.FileServer(http.Dir(Path)))

	if err != nil {
		fmt.Println("File server start error: ", err)
		fmt.Println("Reset port ...")

		Port++
		return run()
	}

	return nil
}
