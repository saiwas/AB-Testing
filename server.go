// package main

// import (
// 	"net/http"

// 	"ab-testing/processors"
// )

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", processors.IndexHandler)

// 	server := &http.Server{
// 		Addr:    "localhost:8080",
// 		Handler: mux,
// 	}

// 	server.ListenAndServe()
// }
