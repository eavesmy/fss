package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/eavesmy/fss/src/service"
)

var Port int

func main() {

	flag.IntVar(&Port, "p", 9989, "-p 80")
	flag.StringVar(&service.Path, "d", "./", "-d ./")
	flag.Parse()

	service.Path, _ = filepath.Abs(service.Path)

	run()
}

func run() error {

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println("File server start error: ", err)
		return nil
	}

	inet := addrs[1]

	fmt.Println("File server start at "+service.Path, strings.Split(inet.String(), "/")[0]+":"+fmt.Sprintf("%d", Port))

	err = http.ListenAndServe(":"+fmt.Sprintf("%d", Port), service.HttpHandler{})

	if err != nil {
		fmt.Println("File server start error: ", err)
		fmt.Println("Reset port ...")

		Port++
		return run()
	}

	return nil
}

func handleLogin() {

}
