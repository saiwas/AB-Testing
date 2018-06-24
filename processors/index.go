package processors

import (
	"fmt"
	"html/template"
	"net/http"
)

type newIndexTemplate struct {
	Title   string
	Context string
}

/*
	Handle the /index route
*/
func IndexHandler(w http.ResponseWriter, request *http.Request) {
	p := newIndexTemplate{
		Title:   "AB Testing Tools",
		Context: "This is the index page",
	}

	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, p)
}
