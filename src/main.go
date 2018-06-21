package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("File server start at http://127.0.0.1:9990")
	err := http.ListenAndServe(":9990", http.FileServer(http.Dir("./")))

	if err != nil {
		fmt.Println("File server start error: ", err)
	}

}
