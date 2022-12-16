package main

import (
	"net/http"

	art "art/pkg"
)

func main() {
	http.HandleFunc("/", art.IndexHandler)
	http.HandleFunc("/ascii", art.Ascii)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}
