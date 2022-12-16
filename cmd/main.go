package main

import (
	"net/http"

	art "art/pkg"
)

func main() {
	http.HandleFunc("/", art.IndexHandler)                                                        // in-app routing systems "/"
	http.HandleFunc("/ascii", art.Ascii)                                                          // in-app routing systems "/ascii"
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))) // file request handler static functions
	http.ListenAndServe(":8080", nil)                                                             // launch web application
}
