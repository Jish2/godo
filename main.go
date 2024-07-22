package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World!\n")
}

func main() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":3000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server Closed\n")
	} else if err != nil {
		fmt.Printf("Error Starting Server: %s\n", err)
		os.Exit(1)
	}
}
