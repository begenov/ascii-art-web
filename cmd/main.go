package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	res, err := template.ParseFiles("templates/index.html", "templates/header.html")
	if err != nil {
		fmt.Println(err)
	}
	res.ExecuteTemplate(w, "index", nil)
}

func ascii_art(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/ascii-art", ascii_art)
	http.ListenAndServe(":8080", nil)
}
