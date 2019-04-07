package main

import (
	"net/http"
	"fmt"
	"log"
)

// Basic logger to help with debugging missing static page errors.
func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s requested %s", r.RemoteAddr, r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {

	// Creating a ServeMux to log messages as well as serve static pages.
	h := http.NewServeMux()

	// Serving static assets.
	fs := http.FileServer(http.Dir("assets/static"))
	h.Handle("/", fs)

	// Connect a basic logger.
	handlerWithLogger := logger(h)

	// Serve at port 8080 on localhost.
	var port = "8080"
	fmt.Println("Serving on port " + port)	
	err := http.ListenAndServe(":" + port, handlerWithLogger)

	// On failure, print the error message instead of existing quietly.
	log.Fatal(err)
}
