package main

import (
	"fmt"
	"log"
	"net/http"
)

type servlet string

func (*servlet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"status\":\"serve\"}")
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"status\":\"handle\"}")
}

func main() {
	s := new(servlet)
	fmt.Println("starting server....")
	http.Handle("/a", s)
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
