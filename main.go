package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome the the picapp site</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<a>Contact me: <a href=\"mailto:adamwoolhether@gmail.com\">adamwoolhether@gmail.com</a></h1>")
}

func lost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Sorry, this page doesn't exist</h1>")
}

func main() {
	r := mux.NewRouter()

	//Creating a custom 404 err response for Gorilla's mux:
	// Because the backed of gorilla's mux NotFoundHandler directs to http.Handler
	// you must implement the http.Handler interface and assign it to NotFoundHandler
	var l http.Handler = http.HandlerFunc(lost)
	r.NotFoundHandler = l

	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe("localhost:3000", r)
}
