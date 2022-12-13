package main

import (
	art "art/pkg"
	"fmt"
	"net/http"
	"text/template"
)

var d struct {
	Input  string `json:"input"`
	Banner string `json:"banner"`
	// Last  string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	res, err := template.ParseFiles("templates/index.html", "templates/header.html")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(1)
	res.ExecuteTemplate(w, "index", nil)
}

func ascii(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	res, err := template.ParseFiles("templates/ascii.html", "templates/header.html")
	if err != nil {
		return
	}
	input := r.FormValue("input")
	inputBanner := r.FormValue("banner")
	// fmt.Println(inputBanner)

	arr := art.Startascii(input, inputBanner)

	d.Input += arr
	d.Banner = inputBanner
	// fmt.Println(d.Input)
	fmt.Println(d)

	res.ExecuteTemplate(w, "ascii", arr)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii", ascii)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)
}
