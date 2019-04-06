package main

import (
	"net/http"
)

func main() {

	//Serving static assets
	fs := http.FileServer(http.Dir("assets/static"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)
}
