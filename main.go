package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") //you can test "text/plain" to see result
	//fmt.Fprint(w, r.URL.Path)
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome the the picapp site</h1>")
	} else if r.URL.Path == "/contact" {
	fmt.Fprint(w, "<a>Contact me: <a href=\"mailto:adamwoolhether@gmail.com\">adamwoolhether@gmail.com</a></h1>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>This page doesn't exist :o</h1><p>Email me if you keep getting this message.</p>")
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()
	router.GET("/hello/:name/mandarin", Hello)
	http.ListenAndServe("localhost:3000", router)
}
