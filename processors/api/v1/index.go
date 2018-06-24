package v1

import (
	"encoding/json"
	"net/http"
)

// Test
type Profile struct {
	Name    string
	Hobbies []string
}

/*
	Handle the /index route
*/
func IndexHandler(w http.ResponseWriter, request *http.Request) {
	profile := Profile{"Johnny", []string{"Coding", "Programming"}}

	js, err := json.Marshal(profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
