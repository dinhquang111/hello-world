package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

var id = uuid.New()

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive request")
	fmt.Fprintf(w, "Hello, World! %s", id)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
