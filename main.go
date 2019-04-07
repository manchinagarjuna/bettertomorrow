package main

import (
	"net/http"
	"fmt"
)

func main() {

	//Serving static assets
	fs := http.FileServer(http.Dir("assets/static"))
	http.Handle("/", fs)

	var port = "8080"
	fmt.Println("Serving on port " + port)
	http.ListenAndServe(":" + port, nil)
}
