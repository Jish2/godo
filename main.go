package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	fmt.Fprintf(w, "ping %d\n", t.Unix())
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", handler)

	err := http.ListenAndServe(":3000", router)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server Closed\n")
	} else if err != nil {
		fmt.Printf("Error Starting Server: %s\n", err)
		os.Exit(1)
	}
}
