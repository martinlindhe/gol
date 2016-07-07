package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	hostAndPort := "localhost:4444"

	fmt.Println("Waiting for connections at http://" + hostAndPort)

	err := http.ListenAndServe(hostAndPort, http.FileServer(http.Dir("./www")))
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
