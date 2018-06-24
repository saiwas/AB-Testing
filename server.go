package main

import (
	"net/http"

	"ab-testing/processors"
	apiV1 "ab-testing/processors/api/v1"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", files))

	mux.HandleFunc("/", processors.IndexHandler)
	mux.HandleFunc("/api/v1/index", apiV1.IndexHandler)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
