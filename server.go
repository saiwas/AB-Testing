package main

import (
	"net/http"

	"ab-testing/processors"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", files))

	mux.HandleFunc("/", processors.IndexHandler)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
