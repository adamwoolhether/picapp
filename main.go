package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome the the picapp site</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, pn httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<a>Contact me: <a href=\"mailto:adamwoolhether@gmail.com\">adamwoolhether@gmail.com</a></h1>")
	fmt.Fprintf(w, "<br><a>You're on the cotact %s page!", pn.ByName("page"))
}

func lost(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>This page is not available</h1>")
}

// This demonstrates the use of httprouter pacakge, and the dynamic routing ability it has.
func main() {
	r := httprouter.New()
	r.GET("/", home)
	r.GET("/contact/:page", contact)
	r.NotFound = http.HandlerFunc(lost)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}
