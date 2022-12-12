package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// /message := []byte("Hello, world!")
	res, err := template.ParseFiles("templates/index.html", "templates/header.html")
	// _, err := w.Write(message)
	// fmt.Println("sadas")
	// log.Fatal(err)
	if err != nil {
		fmt.Println(err)
	}
	res.ExecuteTemplate(w, "index", nil)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}
