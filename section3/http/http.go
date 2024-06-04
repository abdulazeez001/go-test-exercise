package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", HelloHandler)

	port := "8080"
	log.Printf("[pid %d] REST server Listening on port %s", os.Getpid(), port)
	// service connections
	if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
